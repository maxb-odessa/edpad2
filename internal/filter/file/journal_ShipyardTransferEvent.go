package file

import "time"

type ShipyardTransferEvent struct {
	Distance          float64   `mapstructure:"Distance,omitempty"`
	MarketID          int       `mapstructure:"MarketID,omitempty"`
	ShipID            int       `mapstructure:"ShipID,omitempty"`
	ShipMarketID      int       `mapstructure:"ShipMarketID,omitempty"`
	ShipType          string    `mapstructure:"ShipType,omitempty"`
	ShipTypeLocalised string    `mapstructure:"ShipType_Localised,omitempty"`
	System            string    `mapstructure:"System,omitempty"`
	TransferPrice     int       `mapstructure:"TransferPrice,omitempty"`
	TransferTime      int       `mapstructure:"TransferTime,omitempty"`
	Event             string    `mapstructure:"event,omitempty"`
	Timestamp         time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evShipyardTransfer(ev *ShipyardTransferEvent) {
	return
}
