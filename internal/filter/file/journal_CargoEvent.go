package file

import "time"

type CargoEvent struct {
	Count     int `mapstructure:"Count,omitempty"`
	Inventory []struct {
		Count         int    `mapstructure:"Count,omitempty"`
		MissionID     int    `mapstructure:"MissionID,omitempty"`
		Name          string `mapstructure:"Name,omitempty"`
		NameLocalised string `mapstructure:"Name_Localised,omitempty"`
		Stolen        int    `mapstructure:"Stolen,omitempty"`
	} `mapstructure:"Inventory,omitempty"`
	Vessel    string    `mapstructure:"Vessel,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCargo(ev *CargoEvent) {
	return
}
