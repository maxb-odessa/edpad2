package file

import "time"

type SupercruiseExitEvent struct {
	Body          string    `json:"Body,omitempty"`
	BodyID        int       `json:"BodyID,omitempty"`
	BodyType      string    `json:"BodyType,omitempty"`
	Multicrew     bool      `json:"Multicrew,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Taxi          bool      `json:"Taxi,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSupercruiseExit(ev *SupercruiseExitEvent) {
	return
}
