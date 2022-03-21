package file

import "time"

type FSSDiscoveryScanEvent struct {
	BodyCount     int       `json:"BodyCount,omitempty"`
	NonBodyCount  int       `json:"NonBodyCount,omitempty"`
	Progress      float64   `json:"Progress,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	SystemName    string    `json:"SystemName,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFSSDiscoveryScan(ev *FSSDiscoveryScanEvent) {
	return
}
