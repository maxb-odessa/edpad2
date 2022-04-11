package file

import "time"

type DockFighterEvent struct {
	ID        int       `mapstructure:"ID,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evDockFighter(ev *DockFighterEvent) {
	return
}
