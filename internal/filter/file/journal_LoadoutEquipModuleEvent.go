package file

import "time"

type LoadoutEquipModuleEvent struct {
	Class               int           `mapstructure:"Class,omitempty"`
	LoadoutID           int           `mapstructure:"LoadoutID,omitempty"`
	LoadoutName         string        `mapstructure:"LoadoutName,omitempty"`
	ModuleName          string        `mapstructure:"ModuleName,omitempty"`
	ModuleNameLocalised string        `mapstructure:"ModuleName_Localised,omitempty"`
	SlotName            string        `mapstructure:"SlotName,omitempty"`
	SuitID              int           `mapstructure:"SuitID,omitempty"`
	SuitModuleID        int           `mapstructure:"SuitModuleID,omitempty"`
	SuitName            string        `mapstructure:"SuitName,omitempty"`
	SuitNameLocalised   string        `mapstructure:"SuitName_Localised,omitempty"`
	WeaponMods          []interface{} `mapstructure:"WeaponMods,omitempty"`
	Event               string        `mapstructure:"event,omitempty"`
	Timestamp           time.Time     `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evLoadoutEquipModule(ev *LoadoutEquipModuleEvent) {
	return
}
