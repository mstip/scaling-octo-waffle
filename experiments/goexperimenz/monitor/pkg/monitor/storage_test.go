package monitor

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestStorage_GetSystemInfoBySystemId(t *testing.T) {
	type fields struct {
		mux          *sync.RWMutex
		systemInfos  []SystemInfo
		systemStatus map[string][]SystemStatus
	}
	type args struct {
		systemId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SystemInfo
	}{
		{
			name:   "without data",
			fields: fields{mux: &sync.RWMutex{}, systemInfos: []SystemInfo{}, systemStatus: map[string][]SystemStatus{}},
			args:   args{systemId: "does-not-exist"},
			want:   nil,
		},
		{
			name: "with data",
			fields: fields{mux: &sync.RWMutex{}, systemInfos: []SystemInfo{
				{
					SystemID:     "test-id",
					SystemName:   "test-name",
					CustomerID:   1337,
					CustomerName: "test-customer-name",
				},
			}, systemStatus: map[string][]SystemStatus{}},
			args: args{systemId: "test-id"},
			want: &SystemInfo{
				SystemID:     "test-id",
				SystemName:   "test-name",
				CustomerID:   1337,
				CustomerName: "test-customer-name",
			},
		},
		{
			name: "with many data",
			fields: fields{mux: &sync.RWMutex{}, systemInfos: []SystemInfo{
				{
					SystemID:     "test-id",
					SystemName:   "test-name",
					CustomerID:   1337,
					CustomerName: "test-customer-name",
				},
				{
					SystemID:     "test-id2",
					SystemName:   "test-name2",
					CustomerID:   1337,
					CustomerName: "test-customer-name",
				},
				{
					SystemID:     "test-id3",
					SystemName:   "test-name3",
					CustomerID:   1337,
					CustomerName: "test-customer-name",
				},
			}, systemStatus: map[string][]SystemStatus{}},
			args: args{systemId: "test-id"},
			want: &SystemInfo{
				SystemID:     "test-id",
				SystemName:   "test-name",
				CustomerID:   1337,
				CustomerName: "test-customer-name",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				mux:          tt.fields.mux,
				systemInfos:  tt.fields.systemInfos,
				systemStatus: tt.fields.systemStatus,
			}
			if got := s.GetSystemInfoBySystemId(tt.args.systemId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSystemInfoBySystemId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_SaveSystemInfo(t *testing.T) {
	type fields struct {
		mux          *sync.RWMutex
		systemInfos  []SystemInfo
		systemStatus map[string][]SystemStatus
	}
	type args struct {
		info SystemInfo
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		expectedLen int
	}{
		{
			name:   "without data",
			fields: fields{mux: &sync.RWMutex{}, systemInfos: []SystemInfo{}, systemStatus: map[string][]SystemStatus{}},
			args: args{info: SystemInfo{
				SystemID:     "test-id",
				SystemName:   "test-name",
				CustomerID:   1337,
				CustomerName: "test-customer-name",
			}},
			wantErr:     false,
			expectedLen: 1,
		},
		{
			name: "with existing data",
			fields: fields{mux: &sync.RWMutex{}, systemInfos: []SystemInfo{
				{
					SystemID:     "already-existing-test-id",
					SystemName:   "already-existing-test-name",
					CustomerID:   42,
					CustomerName: "already-existing-test-customer-name",
				},
			}, systemStatus: map[string][]SystemStatus{}},
			args: args{info: SystemInfo{
				SystemID:     "test-id",
				SystemName:   "test-name",
				CustomerID:   1337,
				CustomerName: "test-customer-name",
			}},
			wantErr:     false,
			expectedLen: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				mux:          tt.fields.mux,
				systemInfos:  tt.fields.systemInfos,
				systemStatus: tt.fields.systemStatus,
			}
			err := s.SaveSystemInfo(tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveSystemInfo() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(s.systemInfos) != tt.expectedLen {
				t.Errorf("SaveSystemInfo() len = %d,  expectedLen = %d", len(s.systemInfos), tt.expectedLen)
			}
		})
	}
}

func TestStorage_SaveSystemStatus(t *testing.T) {
	type fields struct {
		mux          *sync.RWMutex
		systemInfos  []SystemInfo
		systemStatus map[string][]SystemStatus
	}
	type args struct {
		status SystemStatus
	}
	tests := []struct {
		name                    string
		fields                  fields
		args                    args
		wantErr                 bool
		expectedLen             int
		systemIdToCheck         string
		expectedSystemStatusLen int
	}{
		{
			name:   "without data",
			fields: fields{mux: &sync.RWMutex{}, systemInfos: []SystemInfo{}, systemStatus: map[string][]SystemStatus{}},
			args: args{status: SystemStatus{
				SystemID:         "test-system-id",
				DiskUsagePercent: 10,
				RamUsagePercent:  15,
				LastContact:      time.Now().Unix(),
				Timestamp:        time.Now().Unix(),
			}},
			wantErr:     false,
			expectedLen: 1,
		},
		{
			name: "with data and new system id",
			fields: fields{mux: &sync.RWMutex{}, systemInfos: []SystemInfo{}, systemStatus: map[string][]SystemStatus{
				"test-system-id-2": {
					{
						SystemID:         "test-system-id-2",
						DiskUsagePercent: 10,
						RamUsagePercent:  15,
						LastContact:      time.Now().Unix(),
						Timestamp:        time.Now().Unix(),
					},
				},
			},
			},
			args: args{status: SystemStatus{
				SystemID:         "test-system-id",
				DiskUsagePercent: 10,
				RamUsagePercent:  15,
				LastContact:      time.Now().Unix(),
				Timestamp:        time.Now().Unix(),
			}},
			wantErr:                 false,
			expectedLen:             2,
			systemIdToCheck:         "test-system-id",
			expectedSystemStatusLen: 1,
		},
		{
			name: "with data and existing system id",
			fields: fields{mux: &sync.RWMutex{}, systemInfos: []SystemInfo{}, systemStatus: map[string][]SystemStatus{
				"test-system-id": {
					{
						SystemID:         "test-system-id",
						DiskUsagePercent: 10,
						RamUsagePercent:  15,
						LastContact:      time.Now().Unix(),
						Timestamp:        time.Now().Unix(),
					},
				},
			},
			},
			args: args{status: SystemStatus{
				SystemID:         "test-system-id",
				DiskUsagePercent: 10,
				RamUsagePercent:  15,
				LastContact:      time.Now().Unix(),
				Timestamp:        time.Now().Unix(),
			}},
			wantErr:                 false,
			expectedLen:             1,
			systemIdToCheck:         "test-system-id",
			expectedSystemStatusLen: 2,
		},
		{
			name: "store never more than 10 status",
			fields: fields{mux: &sync.RWMutex{}, systemInfos: []SystemInfo{}, systemStatus: map[string][]SystemStatus{
				"test-system-id": {
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: 1},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: 2},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: 3},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: 4},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: 5},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
				},
			},
			},
			args: args{status: SystemStatus{
				SystemID:         "test-system-id",
				DiskUsagePercent: 10,
				RamUsagePercent:  15,
				LastContact:      time.Now().Unix(),
				Timestamp:        time.Now().Unix(),
			}},
			wantErr:                 false,
			expectedLen:             1,
			systemIdToCheck:         "test-system-id",
			expectedSystemStatusLen: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				mux:          tt.fields.mux,
				systemInfos:  tt.fields.systemInfos,
				systemStatus: tt.fields.systemStatus,
			}
			err := s.SaveSystemStatus(tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveSystemStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(s.systemStatus) != tt.expectedLen {
				t.Errorf("SaveSystemStatus() len = %d,  expectedLen = %d", len(s.systemStatus), tt.expectedLen)
			}

			if tt.systemIdToCheck != "" {
				if _, exists := s.systemStatus[tt.systemIdToCheck]; exists != true {
					t.Errorf("SaveSystemStatus() systemIdToCheck: %s does not exists", tt.systemIdToCheck)
				}

				if len(s.systemStatus[tt.systemIdToCheck]) != tt.expectedSystemStatusLen {
					t.Errorf("SaveSystemStatus() len = %d,  expectedSystemStatusLen = %d", len(s.systemStatus[tt.systemIdToCheck]), tt.expectedSystemStatusLen)
				}
			}
		})
	}
}

func TestStorage_CheckSystemsForLastContactToOld(t *testing.T) {
	type fields struct {
		mux          *sync.RWMutex
		systemInfos  []SystemInfo
		systemStatus map[string][]SystemStatus
	}
	//var empty []SystemInfo
	tests := []struct {
		name    string
		fields  fields
		want    []SystemInfo
		wantErr bool
	}{
		{
			name: "system with 10 ok entries and system with only one not ok",
			fields: fields{mux: &sync.RWMutex{}, systemInfos: []SystemInfo{}, systemStatus: map[string][]SystemStatus{
				// 10 ok entries
				"test-system-id": {
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Unix(), Timestamp: time.Now().Unix()},
				},
				// 120 minutes ago but not 10 entries
				"test-system-id-2": {
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Add(time.Duration(-120) * time.Minute).Unix(), Timestamp: time.Now().Unix()},
				},
			},
			},
		},
		{
			name:    "with no data",
			fields:  fields{mux: &sync.RWMutex{}, systemInfos: []SystemInfo{}, systemStatus: map[string][]SystemStatus{}},
			want:    nil,
			wantErr: false,
		},
		{
			name: "system with 10 not ok entries",
			fields: fields{mux: &sync.RWMutex{}, systemInfos: []SystemInfo{
				{
					SystemID:     "test-system-id",
					SystemName:   "test-name",
					CustomerID:   1337,
					CustomerName: "test-customer-name",
				},
			}, systemStatus: map[string][]SystemStatus{
				// 10 not ok entries
				"test-system-id": {
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Add(time.Duration(-120) * time.Minute).Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Add(time.Duration(-120) * time.Minute).Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Add(time.Duration(-120) * time.Minute).Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Add(time.Duration(-120) * time.Minute).Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Add(time.Duration(-120) * time.Minute).Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Add(time.Duration(-120) * time.Minute).Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Add(time.Duration(-120) * time.Minute).Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Add(time.Duration(-120) * time.Minute).Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Add(time.Duration(-120) * time.Minute).Unix(), Timestamp: time.Now().Unix()},
					{SystemID: "test-system-id", DiskUsagePercent: 10, RamUsagePercent: 15, LastContact: time.Now().Add(time.Duration(-120) * time.Minute).Unix(), Timestamp: time.Now().Unix()},
				},
			},
			},
			wantErr: false,
			want: []SystemInfo{
				{
					SystemID:     "test-system-id",
					SystemName:   "test-name",
					CustomerID:   1337,
					CustomerName: "test-customer-name",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				mux:          tt.fields.mux,
				systemInfos:  tt.fields.systemInfos,
				systemStatus: tt.fields.systemStatus,
			}
			got, err := s.CheckSystemsForLastContactToOld()
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckSystemsForLastContactToOld() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckSystemsForLastContactToOld() got = %v, want %v", got, tt.want)
			}
		})
	}
}
