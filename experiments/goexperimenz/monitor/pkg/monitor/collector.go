package monitor

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type SystemControlMonitorData struct {
	SystemID         string `json:"systemId"`
	SystemName       string `json:"systemName"`
	CustomerID       int    `json:"customerId"`
	CustomerName     string `json:"customerName"`
	Internal         bool   `json:"internal"`
	DiskUsagePercent int    `json:"diskUsagePercent"`
	RamUsagePercent  int    `json:"ramUsagePercent"`
	LastContact      int64  `json:"lastContact"`
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Collector struct {
	storage      *Storage
	httpClient   HTTPClient
	scMonitorUrl string
	scApiUser    string
	scApiPw      string
}

func NewCollector(s *Storage) (*Collector, error) {
	c := &Collector{
		storage:      s,
		httpClient:   &http.Client{},
		scMonitorUrl: os.Getenv("SC_MONITOR_URL"),
		scApiUser:    os.Getenv("SC_API_USER"),
		scApiPw:      os.Getenv("SC_API_PW"),
	}
	return c, nil
}

func (c *Collector) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("collector stopped")
			return
		default:
			err := c.Collect()
			if err != nil {
				log.Println("collector stopped", err)
				return
			}

			time.Sleep(1 * time.Minute)
		}
	}
}

func (c *Collector) Collect() error {
	monitorData, err := c.CollectSystemData()
	if err != nil {
		return err
	}

	err = c.SaveMonitorData(monitorData)
	if err != nil {
		return err
	}
	return nil
}

func (c *Collector) CollectSystemData() ([]SystemControlMonitorData, error) {
	req, err := http.NewRequest(http.MethodGet, c.scMonitorUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("php-auth-user", c.scApiUser)
	req.Header.Set("php-auth-pw", c.scApiPw)
	log.Println("collecting data from: " + c.scMonitorUrl)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("collect failed systemcontrol returned status " + res.Status)
	}

	if res.Body != nil {
		defer func() {
			if err := res.Body.Close(); err != nil {
				log.Println("failed to close body after collecting data", err)
			}
		}()
	}

	if res.Body == nil {
		return nil, errors.New("empty body")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var monitorData []SystemControlMonitorData
	err = json.Unmarshal(body, &monitorData)
	if err != nil {
		return nil, err
	}

	return monitorData, nil
}

func (c *Collector) SaveMonitorData(data []SystemControlMonitorData) error {
	for _, item := range data {
		err := c.storage.SaveSystemInfo(SystemInfo{
			SystemID:     item.SystemID,
			SystemName:   item.SystemName,
			CustomerID:   item.CustomerID,
			CustomerName: item.CustomerName,
		})

		if _, matching := err.(NotUniqueError); err != nil && matching == false {
			return err
		}

		err = c.storage.SaveSystemStatus(SystemStatus{
			SystemID:         item.SystemID,
			DiskUsagePercent: item.DiskUsagePercent,
			RamUsagePercent:  item.RamUsagePercent,
			LastContact:      item.LastContact,
			Timestamp:        time.Now().Unix(),
		})

		if err != nil {
			return err
		}
	}
	return nil
}
