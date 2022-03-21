package file

import "time"

type UndockedEvent struct {
	MarketID    int       `json:"MarketID,omitempty"`
	Multicrew   bool      `json:"Multicrew,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	StationType string    `json:"StationType,omitempty"`
	Taxi        bool      `json:"Taxi,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evUndocked(ev *UndockedEvent) {
	return
}
