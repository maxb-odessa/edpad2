package file

import "time"

type ShipLockerMaterialsEvent struct {
	Components  []interface{} `json:"Components,omitempty"`
	Consumables []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
		OwnerID       int    `json:"OwnerID,omitempty"`
	} `json:"Consumables,omitempty"`
	Data      []interface{} `json:"Data,omitempty"`
	Items     []interface{} `json:"Items,omitempty"`
	Event     string        `json:"event,omitempty"`
	Timestamp time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evShipLockerMaterials(ev *ShipLockerMaterialsEvent) {
	return
}
