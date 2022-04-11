package file

import "time"

type CollectCargoEvent struct {
	Stolen        bool      `mapstructure:"Stolen,omitempty"`
	Type          string    `mapstructure:"Type,omitempty"`
	TypeLocalised string    `mapstructure:"Type_Localised,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCollectCargo(ev *CollectCargoEvent) {
	return
}
