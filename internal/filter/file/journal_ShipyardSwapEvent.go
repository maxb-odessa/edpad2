package file

import "time"

type ShipyardSwapEvent struct {
	MarketID          int       `mapstructure:"MarketID,omitempty"`
	ShipID            int       `mapstructure:"ShipID,omitempty"`
	ShipType          string    `mapstructure:"ShipType,omitempty"`
	ShipTypeLocalised string    `mapstructure:"ShipType_Localised,omitempty"`
	StoreOldShip      string    `mapstructure:"StoreOldShip,omitempty"`
	StoreShipID       int       `mapstructure:"StoreShipID,omitempty"`
	Event             string    `mapstructure:"event,omitempty"`
	Timestamp         time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evShipyardSwap(ev *ShipyardSwapEvent) {
	return
}
