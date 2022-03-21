package file

import "time"

type LoadoutEquipModuleEvent struct {
	Class               int           `json:"Class,omitempty"`
	LoadoutID           int           `json:"LoadoutID,omitempty"`
	LoadoutName         string        `json:"LoadoutName,omitempty"`
	ModuleName          string        `json:"ModuleName,omitempty"`
	ModuleNameLocalised string        `json:"ModuleName_Localised,omitempty"`
	SlotName            string        `json:"SlotName,omitempty"`
	SuitID              int           `json:"SuitID,omitempty"`
	SuitModuleID        int           `json:"SuitModuleID,omitempty"`
	SuitName            string        `json:"SuitName,omitempty"`
	SuitNameLocalised   string        `json:"SuitName_Localised,omitempty"`
	WeaponMods          []interface{} `json:"WeaponMods,omitempty"`
	Event               string        `json:"event,omitempty"`
	Timestamp           time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evLoadoutEquipModule(ev *LoadoutEquipModuleEvent) {
	return
}
