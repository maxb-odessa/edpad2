package file

import "time"

type ShipyardBuyEvent struct {
	MarketID          int       `json:"MarketID,omitempty"`
	SellOldShip       string    `json:"SellOldShip,omitempty"`
	SellPrice         int       `json:"SellPrice,omitempty"`
	SellShipID        int       `json:"SellShipID,omitempty"`
	ShipPrice         int       `json:"ShipPrice,omitempty"`
	ShipType          string    `json:"ShipType,omitempty"`
	ShipTypeLocalised string    `json:"ShipType_Localised,omitempty"`
	StoreOldShip      string    `json:"StoreOldShip,omitempty"`
	StoreShipID       int       `json:"StoreShipID,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipyardBuy(ev *ShipyardBuyEvent) {
	return
}
