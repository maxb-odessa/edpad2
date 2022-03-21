package file

import "time"

type ShipyardSellEvent struct {
	MarketID          int       `json:"MarketID,omitempty"`
	SellShipID        int       `json:"SellShipID,omitempty"`
	ShipPrice         int       `json:"ShipPrice,omitempty"`
	ShipType          string    `json:"ShipType,omitempty"`
	ShipTypeLocalised string    `json:"ShipType_Localised,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipyardSell(ev *ShipyardSellEvent) {
	return
}
