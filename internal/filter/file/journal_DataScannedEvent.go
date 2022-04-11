package file

import "time"

type DataScannedEvent struct {
	Type          string    `mapstructure:"Type,omitempty"`
	TypeLocalised string    `mapstructure:"Type_Localised,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evDataScanned(ev *DataScannedEvent) {
	return
}
