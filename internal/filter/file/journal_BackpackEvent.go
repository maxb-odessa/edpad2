package file

import "time"

type BackpackEvent struct {
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

func (h *handler) evBackpack(ev *BackpackEvent) {
	return
}

type BackPackEvent interface{}

func (h *handler) evBackPack(ev *BackPackEvent) {
	return
}
