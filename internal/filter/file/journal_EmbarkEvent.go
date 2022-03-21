package file

import "time"

type EmbarkEvent struct {
	Body          string    `json:"Body,omitempty"`
	BodyID        int       `json:"BodyID,omitempty"`
	ID            int       `json:"ID,omitempty"`
	MarketID      int       `json:"MarketID,omitempty"`
	Multicrew     bool      `json:"Multicrew,omitempty"`
	OnPlanet      bool      `json:"OnPlanet,omitempty"`
	OnStation     bool      `json:"OnStation,omitempty"`
	Srv           bool      `json:"SRV,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	StationName   string    `json:"StationName,omitempty"`
	StationType   string    `json:"StationType,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Taxi          bool      `json:"Taxi,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evEmbark(ev *EmbarkEvent) {
	return
}
