package file

import "time"

type CreateSuitLoadoutEvent struct {
	LoadoutID   int    `json:"LoadoutID,omitempty"`
	LoadoutName string `json:"LoadoutName,omitempty"`
	Modules     []struct {
		Class               int           `json:"Class,omitempty"`
		ModuleName          string        `json:"ModuleName,omitempty"`
		ModuleNameLocalised string        `json:"ModuleName_Localised,omitempty"`
		SlotName            string        `json:"SlotName,omitempty"`
		SuitModuleID        int           `json:"SuitModuleID,omitempty"`
		WeaponMods          []interface{} `json:"WeaponMods,omitempty"`
	} `json:"Modules,omitempty"`
	SuitID            int           `json:"SuitID,omitempty"`
	SuitMods          []interface{} `json:"SuitMods,omitempty"`
	SuitName          string        `json:"SuitName,omitempty"`
	SuitNameLocalised string        `json:"SuitName_Localised,omitempty"`
	Event             string        `json:"event,omitempty"`
	Timestamp         time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evCreateSuitLoadout(ev *CreateSuitLoadoutEvent) {
	return
}
