package file

import "time"

type DockingCancelledEvent struct {
	MarketID    int       `json:"MarketID,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	StationType string    `json:"StationType,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDockingCancelled(ev *DockingCancelledEvent) {
	return
}
