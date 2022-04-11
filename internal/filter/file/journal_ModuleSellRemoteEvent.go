package file

import "time"

type ModuleSellRemoteEvent struct {
	SellItem          string    `mapstructure:"SellItem,omitempty"`
	SellItemLocalised string    `mapstructure:"SellItem_Localised,omitempty"`
	SellPrice         int       `mapstructure:"SellPrice,omitempty"`
	ServerID          int       `mapstructure:"ServerId,omitempty"`
	Ship              string    `mapstructure:"Ship,omitempty"`
	ShipID            int       `mapstructure:"ShipID,omitempty"`
	StorageSlot       int       `mapstructure:"StorageSlot,omitempty"`
	Event             string    `mapstructure:"event,omitempty"`
	Timestamp         time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evModuleSellRemote(ev *ModuleSellRemoteEvent) {
	return
}
