package file

import "time"

type PowerplayJoinEvent struct {
	Power     string    `mapstructure:"Power,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evPowerplayJoin(ev *PowerplayJoinEvent) {
	return
}
