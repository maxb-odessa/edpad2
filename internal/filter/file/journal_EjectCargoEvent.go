package file

import "time"

type EjectCargoEvent struct {
	Abandoned     bool      `mapstructure:"Abandoned,omitempty"`
	Count         int       `mapstructure:"Count,omitempty"`
	Type          string    `mapstructure:"Type,omitempty"`
	TypeLocalised string    `mapstructure:"Type_Localised,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evEjectCargo(ev *EjectCargoEvent) {
	return
}
