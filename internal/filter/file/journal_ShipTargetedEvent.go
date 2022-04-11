package file

import "time"

type ShipTargetedEvent struct {
	Bounty             int       `mapstructure:"Bounty,omitempty"`
	Faction            string    `mapstructure:"Faction,omitempty"`
	HullHealth         float64   `mapstructure:"HullHealth,omitempty"`
	LegalStatus        string    `mapstructure:"LegalStatus,omitempty"`
	PilotName          string    `mapstructure:"PilotName,omitempty"`
	PilotNameLocalised string    `mapstructure:"PilotName_Localised,omitempty"`
	PilotRank          string    `mapstructure:"PilotRank,omitempty"`
	ScanStage          int       `mapstructure:"ScanStage,omitempty"`
	ShieldHealth       float64   `mapstructure:"ShieldHealth,omitempty"`
	Ship               string    `mapstructure:"Ship,omitempty"`
	ShipLocalised      string    `mapstructure:"Ship_Localised,omitempty"`
	SquadronID         string    `mapstructure:"SquadronID,omitempty"`
	TargetLocked       bool      `mapstructure:"TargetLocked,omitempty"`
	Event              string    `mapstructure:"event,omitempty"`
	Timestamp          time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evShipTargeted(ev *ShipTargetedEvent) {
	return
}
