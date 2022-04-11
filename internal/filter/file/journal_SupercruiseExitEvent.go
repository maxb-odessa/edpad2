package file

import "time"

type SupercruiseExitEvent struct {
	Body          string    `mapstructure:"Body,omitempty"`
	BodyID        int       `mapstructure:"BodyID,omitempty"`
	BodyType      string    `mapstructure:"BodyType,omitempty"`
	Multicrew     bool      `mapstructure:"Multicrew,omitempty"`
	StarSystem    string    `mapstructure:"StarSystem,omitempty"`
	SystemAddress int       `mapstructure:"SystemAddress,omitempty"`
	Taxi          bool      `mapstructure:"Taxi,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evSupercruiseExit(ev *SupercruiseExitEvent) {
	return
}
