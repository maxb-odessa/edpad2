package file

import "time"

type CreateSuitLoadoutEvent struct {
	LoadoutID   int    `mapstructure:"LoadoutID,omitempty"`
	LoadoutName string `mapstructure:"LoadoutName,omitempty"`
	Modules     []struct {
		Class               int           `mapstructure:"Class,omitempty"`
		ModuleName          string        `mapstructure:"ModuleName,omitempty"`
		ModuleNameLocalised string        `mapstructure:"ModuleName_Localised,omitempty"`
		SlotName            string        `mapstructure:"SlotName,omitempty"`
		SuitModuleID        int           `mapstructure:"SuitModuleID,omitempty"`
		WeaponMods          []interface{} `mapstructure:"WeaponMods,omitempty"`
	} `mapstructure:"Modules,omitempty"`
	SuitID            int           `mapstructure:"SuitID,omitempty"`
	SuitMods          []interface{} `mapstructure:"SuitMods,omitempty"`
	SuitName          string        `mapstructure:"SuitName,omitempty"`
	SuitNameLocalised string        `mapstructure:"SuitName_Localised,omitempty"`
	Event             string        `mapstructure:"event,omitempty"`
	Timestamp         time.Time     `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCreateSuitLoadout(ev *CreateSuitLoadoutEvent) {
	return
}
