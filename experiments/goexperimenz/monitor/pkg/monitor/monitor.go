package monitor

import (
	"context"
	"fmt"
	"log"
	"os"
	"platform/pkg/pushover"
	"time"
)

type Monitor struct {
	storage    *Storage
	collector  *Collector
	pushover   *pushover.Pushover
	alertDelay int
}

func NewMonitor() (*Monitor, error) {
	s, err := NewStorage()
	if err != nil {
		return nil, err
	}

	c, err := NewCollector(s)
	if err != nil {
		return nil, err
	}

	p := pushover.NewPushover(os.Getenv("PUSHOVER_TOKEN"), os.Getenv("PUSHOVER_USER"))

	return &Monitor{
		storage:    s,
		collector:  c,
		pushover:   p,
		alertDelay: 10,
	}, nil
}

func (m *Monitor) CollectorAnalyzeAlert() error {
	if err := m.collector.Collect(); err != nil {
		return err
	}
	//Analyze
	log.Println("analyze")
	lastContactToOldSystems, err := m.storage.CheckSystemsForLastContactToOld()
	if err != nil {
		return err
	}

	// Alert
	var alertMsg string
	hasError := false
	log.Println("alert")
	if len(lastContactToOldSystems) > 0 {
		hasError = true
		alertMsg += "Last Contact To Old\n"
		alertMsg += "==========\n"
		for _, item := range lastContactToOldSystems {
			alertMsg += fmt.Sprintf("Customer %s System: %s \n", item.CustomerName, item.SystemName)
		}
	}

	if hasError {
		m.alertDelay -= 1
		if m.alertDelay == 0 {
			if err := m.pushover.SendMessage(alertMsg); err != nil {
				return err
			}
			m.alertDelay = 10
		}
	}

	return nil
}

func (m *Monitor) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			log.Println("monitor stopped -> done")
			return nil
		default:
			if err := m.CollectorAnalyzeAlert(); err != nil {
				log.Println("monitor stopped", err)
				return err
			}

			//time.Sleep(1 * time.Minute)
			time.Sleep(3 * time.Second)
		}
	}
}
