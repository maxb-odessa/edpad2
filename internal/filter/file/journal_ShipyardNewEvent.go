package file

import "time"

type ShipyardNewEvent struct {
	NewShipID         int       `mapstructure:"NewShipID,omitempty"`
	ShipType          string    `mapstructure:"ShipType,omitempty"`
	ShipTypeLocalised string    `mapstructure:"ShipType_Localised,omitempty"`
	Event             string    `mapstructure:"event,omitempty"`
	Timestamp         time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evShipyardNew(ev *ShipyardNewEvent) {
	return
}
