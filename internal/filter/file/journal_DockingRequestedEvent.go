package file

import "time"

type DockingRequestedEvent struct {
	MarketID    int       `mapstructure:"MarketID,omitempty"`
	StationName string    `mapstructure:"StationName,omitempty"`
	StationType string    `mapstructure:"StationType,omitempty"`
	Event       string    `mapstructure:"event,omitempty"`
	Timestamp   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evDockingRequested(ev *DockingRequestedEvent) {
	return
}
