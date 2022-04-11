package file

import "time"

type LeaveBodyEvent struct {
	Body          string    `mapstructure:"Body,omitempty"`
	BodyID        int       `mapstructure:"BodyID,omitempty"`
	StarSystem    string    `mapstructure:"StarSystem,omitempty"`
	SystemAddress int       `mapstructure:"SystemAddress,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evLeaveBody(ev *LeaveBodyEvent) {
	return
}
