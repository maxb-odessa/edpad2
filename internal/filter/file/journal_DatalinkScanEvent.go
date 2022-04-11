package file

import "time"

type DatalinkScanEvent struct {
	Message          string    `mapstructure:"Message,omitempty"`
	MessageLocalised string    `mapstructure:"Message_Localised,omitempty"`
	Event            string    `mapstructure:"event,omitempty"`
	Timestamp        time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evDatalinkScan(ev *DatalinkScanEvent) {
	return
}
