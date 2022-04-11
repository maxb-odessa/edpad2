package file

import "time"

type ShieldStateEvent struct {
	ShieldsUp bool      `mapstructure:"ShieldsUp,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evShieldState(ev *ShieldStateEvent) {
	return
}
