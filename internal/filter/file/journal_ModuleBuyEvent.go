package file

import "time"

type ModuleBuyEvent struct {
	BuyItem             string    `mapstructure:"BuyItem,omitempty"`
	BuyItemLocalised    string    `mapstructure:"BuyItem_Localised,omitempty"`
	BuyPrice            int       `mapstructure:"BuyPrice,omitempty"`
	MarketID            int       `mapstructure:"MarketID,omitempty"`
	SellItem            string    `mapstructure:"SellItem,omitempty"`
	SellItemLocalised   string    `mapstructure:"SellItem_Localised,omitempty"`
	SellPrice           int       `mapstructure:"SellPrice,omitempty"`
	Ship                string    `mapstructure:"Ship,omitempty"`
	ShipID              int       `mapstructure:"ShipID,omitempty"`
	Slot                string    `mapstructure:"Slot,omitempty"`
	StoredItem          string    `mapstructure:"StoredItem,omitempty"`
	StoredItemLocalised string    `mapstructure:"StoredItem_Localised,omitempty"`
	Event               string    `mapstructure:"event,omitempty"`
	Timestamp           time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evModuleBuy(ev *ModuleBuyEvent) {
	return
}
