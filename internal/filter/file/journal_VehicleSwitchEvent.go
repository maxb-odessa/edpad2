package file

import "time"

type VehicleSwitchEvent struct {
	To        string    `json:"To,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evVehicleSwitch(ev *VehicleSwitchEvent) {
	return
}
