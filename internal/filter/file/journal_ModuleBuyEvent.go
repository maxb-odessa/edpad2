package file

import "time"

type ModuleBuyEvent struct {
	BuyItem             string    `json:"BuyItem,omitempty"`
	BuyItemLocalised    string    `json:"BuyItem_Localised,omitempty"`
	BuyPrice            int       `json:"BuyPrice,omitempty"`
	MarketID            int       `json:"MarketID,omitempty"`
	SellItem            string    `json:"SellItem,omitempty"`
	SellItemLocalised   string    `json:"SellItem_Localised,omitempty"`
	SellPrice           int       `json:"SellPrice,omitempty"`
	Ship                string    `json:"Ship,omitempty"`
	ShipID              int       `json:"ShipID,omitempty"`
	Slot                string    `json:"Slot,omitempty"`
	StoredItem          string    `json:"StoredItem,omitempty"`
	StoredItemLocalised string    `json:"StoredItem_Localised,omitempty"`
	Event               string    `json:"event,omitempty"`
	Timestamp           time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evModuleBuy(ev *ModuleBuyEvent) {
	return
}
