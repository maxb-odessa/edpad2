package file

import "time"

type DatalinkScanEvent struct {
	Message          string    `json:"Message,omitempty"`
	MessageLocalised string    `json:"Message_Localised,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDatalinkScan(ev *DatalinkScanEvent) {
	return
}
