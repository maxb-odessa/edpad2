package file

import "time"

type DockingDeniedEvent struct {
	MarketID    int       `mapstructure:"MarketID,omitempty"`
	Reason      string    `mapstructure:"Reason,omitempty"`
	StationName string    `mapstructure:"StationName,omitempty"`
	StationType string    `mapstructure:"StationType,omitempty"`
	Event       string    `mapstructure:"event,omitempty"`
	Timestamp   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evDockingDenied(ev *DockingDeniedEvent) {
	return
}
