package file

import "time"

type ShipyardNewEvent struct {
	NewShipID         int       `json:"NewShipID,omitempty"`
	ShipType          string    `json:"ShipType,omitempty"`
	ShipTypeLocalised string    `json:"ShipType_Localised,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipyardNew(ev *ShipyardNewEvent) {
	return
}
