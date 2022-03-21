package file

import "time"

type DockingGrantedEvent struct {
	LandingPad  int       `json:"LandingPad,omitempty"`
	MarketID    int       `json:"MarketID,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	StationType string    `json:"StationType,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDockingGranted(ev *DockingGrantedEvent) {
	return
}
