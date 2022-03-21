package file

import "time"

type BuyWeaponEvent struct {
	Class         int           `json:"Class,omitempty"`
	Name          string        `json:"Name,omitempty"`
	NameLocalised string        `json:"Name_Localised,omitempty"`
	Price         int           `json:"Price,omitempty"`
	SuitModuleID  int           `json:"SuitModuleID,omitempty"`
	WeaponMods    []interface{} `json:"WeaponMods,omitempty"`
	Event         string        `json:"event,omitempty"`
	Timestamp     time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evBuyWeapon(ev *BuyWeaponEvent) {
	return
}
