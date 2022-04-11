package file

import "time"

type ModuleSellEvent struct {
	MarketID          int       `mapstructure:"MarketID,omitempty"`
	SellItem          string    `mapstructure:"SellItem,omitempty"`
	SellItemLocalised string    `mapstructure:"SellItem_Localised,omitempty"`
	SellPrice         int       `mapstructure:"SellPrice,omitempty"`
	Ship              string    `mapstructure:"Ship,omitempty"`
	ShipID            int       `mapstructure:"ShipID,omitempty"`
	Slot              string    `mapstructure:"Slot,omitempty"`
	Event             string    `mapstructure:"event,omitempty"`
	Timestamp         time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evModuleSell(ev *ModuleSellEvent) {
	return
}
