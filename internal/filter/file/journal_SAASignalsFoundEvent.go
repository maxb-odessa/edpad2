package file

import "time"

type SAASignalsFoundEvent struct {
	BodyID   int    `json:"BodyID,omitempty"`
	BodyName string `json:"BodyName,omitempty"`
	Signals  []struct {
		Count         int    `json:"Count,omitempty"`
		Type          string `json:"Type,omitempty"`
		TypeLocalised string `json:"Type_Localised,omitempty"`
	} `json:"Signals,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSAASignalsFound(ev *SAASignalsFoundEvent) {
	return
}
