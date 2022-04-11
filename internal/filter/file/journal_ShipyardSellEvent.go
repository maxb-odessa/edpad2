package file

import "time"

type ShipyardSellEvent struct {
	MarketID          int       `mapstructure:"MarketID,omitempty"`
	SellShipID        int       `mapstructure:"SellShipID,omitempty"`
	ShipPrice         int       `mapstructure:"ShipPrice,omitempty"`
	ShipType          string    `mapstructure:"ShipType,omitempty"`
	ShipTypeLocalised string    `mapstructure:"ShipType_Localised,omitempty"`
	Event             string    `mapstructure:"event,omitempty"`
	Timestamp         time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evShipyardSell(ev *ShipyardSellEvent) {
	return
}
