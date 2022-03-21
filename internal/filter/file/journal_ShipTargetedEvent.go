package file

import "time"

type ShipTargetedEvent struct {
	Bounty             int       `json:"Bounty,omitempty"`
	Faction            string    `json:"Faction,omitempty"`
	HullHealth         float64   `json:"HullHealth,omitempty"`
	LegalStatus        string    `json:"LegalStatus,omitempty"`
	PilotName          string    `json:"PilotName,omitempty"`
	PilotNameLocalised string    `json:"PilotName_Localised,omitempty"`
	PilotRank          string    `json:"PilotRank,omitempty"`
	ScanStage          int       `json:"ScanStage,omitempty"`
	ShieldHealth       float64   `json:"ShieldHealth,omitempty"`
	Ship               string    `json:"Ship,omitempty"`
	ShipLocalised      string    `json:"Ship_Localised,omitempty"`
	SquadronID         string    `json:"SquadronID,omitempty"`
	TargetLocked       bool      `json:"TargetLocked,omitempty"`
	Event              string    `json:"event,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipTargeted(ev *ShipTargetedEvent) {
	return
}
