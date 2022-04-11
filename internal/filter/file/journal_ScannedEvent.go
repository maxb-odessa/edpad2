package file

import "time"

type ScannedEvent struct {
	ScanType  string    `mapstructure:"ScanType,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evScanned(ev *ScannedEvent) {
	return
}
