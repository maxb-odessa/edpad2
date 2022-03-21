package file

import "time"

type PowerplayJoinEvent struct {
	Power     string    `json:"Power,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPowerplayJoin(ev *PowerplayJoinEvent) {
	return
}
