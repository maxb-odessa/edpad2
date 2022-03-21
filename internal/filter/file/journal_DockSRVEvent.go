package file

import "time"

type DockSRVEvent struct {
	ID        int       `json:"ID,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDockSRV(ev *DockSRVEvent) {
	return
}
