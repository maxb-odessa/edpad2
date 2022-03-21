package file

import "time"

type ShipyardTransferEvent struct {
	Distance          float64   `json:"Distance,omitempty"`
	MarketID          int       `json:"MarketID,omitempty"`
	ShipID            int       `json:"ShipID,omitempty"`
	ShipMarketID      int       `json:"ShipMarketID,omitempty"`
	ShipType          string    `json:"ShipType,omitempty"`
	ShipTypeLocalised string    `json:"ShipType_Localised,omitempty"`
	System            string    `json:"System,omitempty"`
	TransferPrice     int       `json:"TransferPrice,omitempty"`
	TransferTime      int       `json:"TransferTime,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipyardTransfer(ev *ShipyardTransferEvent) {
	return
}
