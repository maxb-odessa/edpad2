package file

import "time"

type SupercruiseEntryEvent struct {
	Multicrew     bool      `json:"Multicrew,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Taxi          bool      `json:"Taxi,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSupercruiseEntry(ev *SupercruiseEntryEvent) {
	return
}
