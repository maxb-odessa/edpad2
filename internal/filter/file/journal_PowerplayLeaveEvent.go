package file

import "time"

type PowerplayLeaveEvent struct {
	Power     string    `json:"Power,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPowerplayLeave(ev *PowerplayLeaveEvent) {
	return
}
