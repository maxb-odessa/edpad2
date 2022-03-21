package file

import "time"

type DockFighterEvent struct {
	ID        int       `json:"ID,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDockFighter(ev *DockFighterEvent) {
	return
}
