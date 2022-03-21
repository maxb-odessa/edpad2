package file

import "time"

type ScannedEvent struct {
	ScanType  string    `json:"ScanType,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evScanned(ev *ScannedEvent) {
	return
}
