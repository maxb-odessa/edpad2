package file

import "time"

type ShipyardBuyEvent struct {
	MarketID          int       `mapstructure:"MarketID,omitempty"`
	SellOldShip       string    `mapstructure:"SellOldShip,omitempty"`
	SellPrice         int       `mapstructure:"SellPrice,omitempty"`
	SellShipID        int       `mapstructure:"SellShipID,omitempty"`
	ShipPrice         int       `mapstructure:"ShipPrice,omitempty"`
	ShipType          string    `mapstructure:"ShipType,omitempty"`
	ShipTypeLocalised string    `mapstructure:"ShipType_Localised,omitempty"`
	StoreOldShip      string    `mapstructure:"StoreOldShip,omitempty"`
	StoreShipID       int       `mapstructure:"StoreShipID,omitempty"`
	Event             string    `mapstructure:"event,omitempty"`
	Timestamp         time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evShipyardBuy(ev *ShipyardBuyEvent) {
	return
}
