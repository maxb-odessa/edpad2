package file

import "time"

type CargoEvent struct {
	Count     int `json:"Count,omitempty"`
	Inventory []struct {
		Count         int    `json:"Count,omitempty"`
		MissionID     int    `json:"MissionID,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
		Stolen        int    `json:"Stolen,omitempty"`
	} `json:"Inventory,omitempty"`
	Vessel    string    `json:"Vessel,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCargo(ev *CargoEvent) {
	return
}
