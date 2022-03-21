package file

import "time"

type ModuleSellRemoteEvent struct {
	SellItem          string    `json:"SellItem,omitempty"`
	SellItemLocalised string    `json:"SellItem_Localised,omitempty"`
	SellPrice         int       `json:"SellPrice,omitempty"`
	ServerID          int       `json:"ServerId,omitempty"`
	Ship              string    `json:"Ship,omitempty"`
	ShipID            int       `json:"ShipID,omitempty"`
	StorageSlot       int       `json:"StorageSlot,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evModuleSellRemote(ev *ModuleSellRemoteEvent) {
	return
}
