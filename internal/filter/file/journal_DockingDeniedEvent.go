package file

import "time"

type DockingDeniedEvent struct {
	MarketID    int       `json:"MarketID,omitempty"`
	Reason      string    `json:"Reason,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	StationType string    `json:"StationType,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDockingDenied(ev *DockingDeniedEvent) {
	return
}
