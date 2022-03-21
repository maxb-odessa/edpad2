package file

import "time"

type MarketEvent struct {
	CarrierDockingAccess string    `json:"CarrierDockingAccess,omitempty"`
	MarketID             int       `json:"MarketID,omitempty"`
	StarSystem           string    `json:"StarSystem,omitempty"`
	StationName          string    `json:"StationName,omitempty"`
	StationType          string    `json:"StationType,omitempty"`
	Event                string    `json:"event,omitempty"`
	Timestamp            time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMarket(ev *MarketEvent) {
	return
}
