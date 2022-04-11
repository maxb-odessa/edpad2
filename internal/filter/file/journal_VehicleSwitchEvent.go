package file

import "time"

type VehicleSwitchEvent struct {
	To        string    `mapstructure:"To,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evVehicleSwitch(ev *VehicleSwitchEvent) {
	return
}
