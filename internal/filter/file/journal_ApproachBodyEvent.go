package file

import "time"

type ApproachBodyEvent struct {
	Body          string    `mapstructure:"Body,omitempty"`
	BodyID        int       `mapstructure:"BodyID,omitempty"`
	StarSystem    string    `mapstructure:"StarSystem,omitempty"`
	SystemAddress int       `mapstructure:"SystemAddress,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evApproachBody(ev *ApproachBodyEvent) {
	return
}
