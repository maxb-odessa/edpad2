package file

import "time"

type ShipyardSwapEvent struct {
	MarketID          int       `json:"MarketID,omitempty"`
	ShipID            int       `json:"ShipID,omitempty"`
	ShipType          string    `json:"ShipType,omitempty"`
	ShipTypeLocalised string    `json:"ShipType_Localised,omitempty"`
	StoreOldShip      string    `json:"StoreOldShip,omitempty"`
	StoreShipID       int       `json:"StoreShipID,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipyardSwap(ev *ShipyardSwapEvent) {
	return
}
