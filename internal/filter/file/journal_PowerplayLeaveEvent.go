package file

import "time"

type PowerplayLeaveEvent struct {
	Power     string    `mapstructure:"Power,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evPowerplayLeave(ev *PowerplayLeaveEvent) {
	return
}
