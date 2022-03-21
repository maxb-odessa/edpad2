package file

import "time"

type CollectCargoEvent struct {
	Stolen        bool      `json:"Stolen,omitempty"`
	Type          string    `json:"Type,omitempty"`
	TypeLocalised string    `json:"Type_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCollectCargo(ev *CollectCargoEvent) {
	return
}
