package file

import "time"

type DataScannedEvent struct {
	Type          string    `json:"Type,omitempty"`
	TypeLocalised string    `json:"Type_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDataScanned(ev *DataScannedEvent) {
	return
}
