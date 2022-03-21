package file

import "time"

type EjectCargoEvent struct {
	Abandoned     bool      `json:"Abandoned,omitempty"`
	Count         int       `json:"Count,omitempty"`
	Type          string    `json:"Type,omitempty"`
	TypeLocalised string    `json:"Type_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evEjectCargo(ev *EjectCargoEvent) {
	return
}
