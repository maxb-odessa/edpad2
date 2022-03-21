package file

import "time"

type ApproachSettlementEvent struct {
	BodyID        int       `json:"BodyID,omitempty"`
	BodyName      string    `json:"BodyName,omitempty"`
	Latitude      float64   `json:"Latitude,omitempty"`
	Longitude     float64   `json:"Longitude,omitempty"`
	MarketID      int       `json:"MarketID,omitempty"`
	Name          string    `json:"Name,omitempty"`
	NameLocalised string    `json:"Name_Localised,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evApproachSettlement(ev *ApproachSettlementEvent) {
	return
}
