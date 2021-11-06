package monitor

import (
	"log"
	"sync"
	"time"
)

const MaxStatus = 10

//SystemInfo holds infos about a systemcontrol system.
type SystemInfo struct {
	SystemID     string // uuid unique in sc
	SystemName   string // user friendly name of the system
	CustomerID   int    // customer id the system belongs to
	CustomerName string // user friendly customer name
}

type SystemStatus struct {
	SystemID         string
	DiskUsagePercent int
	RamUsagePercent  int
	LastContact      int64
	Timestamp        int64
}

type NotUniqueError struct {
	err string
}

func (n NotUniqueError) Error() string {
	return n.err
}

type Storage struct {
	mux          *sync.RWMutex
	systemInfos  []SystemInfo
	systemStatus map[string][]SystemStatus
}

func NewStorage() (*Storage, error) {
	return &Storage{
		mux:          &sync.RWMutex{},
		systemStatus: map[string][]SystemStatus{},
	}, nil
}

func (s *Storage) SaveSystemInfo(info SystemInfo) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	if info := s.GetSystemInfoBySystemId(info.SystemID); info != nil {
		return NotUniqueError{err: "a system info with id: " + info.SystemID + " already exists"}
	}
	s.systemInfos = append(s.systemInfos, info)
	return nil
}

func (s *Storage) GetSystemInfoBySystemId(systemId string) *SystemInfo {
	for _, v := range s.systemInfos {
		if v.SystemID == systemId {
			return &v
		}
	}
	return nil
}

func (s *Storage) SaveSystemStatus(status SystemStatus) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.systemStatus[status.SystemID] = append(s.systemStatus[status.SystemID], status)
	// keep only the 10 recent status
	if len(s.systemStatus[status.SystemID]) > MaxStatus {
		s.systemStatus[status.SystemID] = s.systemStatus[status.SystemID][1:len(s.systemStatus[status.SystemID])]
	}
	return nil
}

//CheckSystemsForLastContactToOld return systems where the last contact is in every entry older than an hour
func (s *Storage) CheckSystemsForLastContactToOld() ([]SystemInfo, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	var alertSystems []SystemInfo
	for systemId, statusItems := range s.systemStatus {
		if len(statusItems) < MaxStatus {
			continue
		}
		alertCount := 0
		for _, item := range statusItems {
			if item.LastContact < time.Now().Add(time.Duration(-60)*time.Minute).Unix() {
				alertCount += 1
			}
		}
		if alertCount >= MaxStatus {
			info := s.GetSystemInfoBySystemId(systemId)
			if info != nil {
				alertSystems = append(alertSystems, *info)
			} else {
				log.Fatalf("systemid " + systemId + " was removed")
			}
		}
	}
	return alertSystems, nil
}
