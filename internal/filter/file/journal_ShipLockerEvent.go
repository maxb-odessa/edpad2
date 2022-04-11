package file

import "time"

type ShipLockerEvent struct {
	Components  []interface{} `mapstructure:"Components,omitempty"`
	Consumables []struct {
		Count         int    `mapstructure:"Count,omitempty"`
		Name          string `mapstructure:"Name,omitempty"`
		NameLocalised string `mapstructure:"Name_Localised,omitempty"`
		OwnerID       int    `mapstructure:"OwnerID,omitempty"`
	} `mapstructure:"Consumables,omitempty"`
	Data      []interface{} `mapstructure:"Data,omitempty"`
	Items     []interface{} `mapstructure:"Items,omitempty"`
	Event     string        `mapstructure:"event,omitempty"`
	Timestamp time.Time     `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evShipLocker(ev *ShipLockerEvent) {
	return
}
