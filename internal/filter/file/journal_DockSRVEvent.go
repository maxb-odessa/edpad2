package file

import "time"

type DockSRVEvent struct {
	ID        int       `mapstructure:"ID,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evDockSRV(ev *DockSRVEvent) {
	return
}
