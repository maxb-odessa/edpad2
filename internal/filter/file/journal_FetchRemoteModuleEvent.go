package file

import "time"

type FetchRemoteModuleEvent struct {
	ServerID            int       `mapstructure:"ServerId,omitempty"`
	Ship                string    `mapstructure:"Ship,omitempty"`
	ShipID              int       `mapstructure:"ShipID,omitempty"`
	StorageSlot         int       `mapstructure:"StorageSlot,omitempty"`
	StoredItem          string    `mapstructure:"StoredItem,omitempty"`
	StoredItemLocalised string    `mapstructure:"StoredItem_Localised,omitempty"`
	TransferCost        int       `mapstructure:"TransferCost,omitempty"`
	TransferTime        int       `mapstructure:"TransferTime,omitempty"`
	Event               string    `mapstructure:"event,omitempty"`
	Timestamp           time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evFetchRemoteModule(ev *FetchRemoteModuleEvent) {
	return
}
