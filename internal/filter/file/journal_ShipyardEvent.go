package file

import "time"

type ShipyardEvent struct {
	MarketID    int       `json:"MarketID,omitempty"`
	StarSystem  string    `json:"StarSystem,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipyard(ev *ShipyardEvent) {
	return
}
