package file

import "time"

type DockingRequestedEvent struct {
	MarketID    int       `json:"MarketID,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	StationType string    `json:"StationType,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDockingRequested(ev *DockingRequestedEvent) {
	return
}
