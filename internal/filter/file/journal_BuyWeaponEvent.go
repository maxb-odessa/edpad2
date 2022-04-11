package file

import "time"

type BuyWeaponEvent struct {
	Class         int           `mapstructure:"Class,omitempty"`
	Name          string        `mapstructure:"Name,omitempty"`
	NameLocalised string        `mapstructure:"Name_Localised,omitempty"`
	Price         int           `mapstructure:"Price,omitempty"`
	SuitModuleID  int           `mapstructure:"SuitModuleID,omitempty"`
	WeaponMods    []interface{} `mapstructure:"WeaponMods,omitempty"`
	Event         string        `mapstructure:"event,omitempty"`
	Timestamp     time.Time     `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evBuyWeapon(ev *BuyWeaponEvent) {
	return
}
