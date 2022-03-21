package file

import "time"

type ShieldStateEvent struct {
	ShieldsUp bool      `json:"ShieldsUp,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShieldState(ev *ShieldStateEvent) {
	return
}
