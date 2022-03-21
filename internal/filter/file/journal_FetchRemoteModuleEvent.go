package file

import "time"

type FetchRemoteModuleEvent struct {
	ServerID            int       `json:"ServerId,omitempty"`
	Ship                string    `json:"Ship,omitempty"`
	ShipID              int       `json:"ShipID,omitempty"`
	StorageSlot         int       `json:"StorageSlot,omitempty"`
	StoredItem          string    `json:"StoredItem,omitempty"`
	StoredItemLocalised string    `json:"StoredItem_Localised,omitempty"`
	TransferCost        int       `json:"TransferCost,omitempty"`
	TransferTime        int       `json:"TransferTime,omitempty"`
	Event               string    `json:"event,omitempty"`
	Timestamp           time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFetchRemoteModule(ev *FetchRemoteModuleEvent) {
	return
}
