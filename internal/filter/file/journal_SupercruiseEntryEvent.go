package file

import "time"

type SupercruiseEntryEvent struct {
	Multicrew     bool      `mapstructure:"Multicrew,omitempty"`
	StarSystem    string    `mapstructure:"StarSystem,omitempty"`
	SystemAddress int       `mapstructure:"SystemAddress,omitempty"`
	Taxi          bool      `mapstructure:"Taxi,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evSupercruiseEntry(ev *SupercruiseEntryEvent) {
	return
}
