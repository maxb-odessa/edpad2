package file

import "time"

type ModuleSellEvent struct {
	MarketID          int       `json:"MarketID,omitempty"`
	SellItem          string    `json:"SellItem,omitempty"`
	SellItemLocalised string    `json:"SellItem_Localised,omitempty"`
	SellPrice         int       `json:"SellPrice,omitempty"`
	Ship              string    `json:"Ship,omitempty"`
	ShipID            int       `json:"ShipID,omitempty"`
	Slot              string    `json:"Slot,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evModuleSell(ev *ModuleSellEvent) {
	return
}
