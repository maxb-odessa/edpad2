package file

import "time"

type MissionAcceptedEvent struct {
	Commodity           string    `json:"Commodity,omitempty"`
	CommodityLocalised  string    `json:"Commodity_Localised,omitempty"`
	Count               int       `json:"Count,omitempty"`
	DestinationStation  string    `json:"DestinationStation,omitempty"`
	DestinationSystem   string    `json:"DestinationSystem,omitempty"`
	Donation            string    `json:"Donation,omitempty"`
	Expiry              time.Time `json:"Expiry,omitempty"`
	Faction             string    `json:"Faction,omitempty"`
	Influence           string    `json:"Influence,omitempty"`
	LocalisedName       string    `json:"LocalisedName,omitempty"`
	MissionID           int       `json:"MissionID,omitempty"`
	Name                string    `json:"Name,omitempty"`
	PassengerCount      int       `json:"PassengerCount,omitempty"`
	PassengerType       string    `json:"PassengerType,omitempty"`
	PassengerViPs       bool      `json:"PassengerVIPs,omitempty"`
	PassengerWanted     bool      `json:"PassengerWanted,omitempty"`
	Reputation          string    `json:"Reputation,omitempty"`
	Reward              int       `json:"Reward,omitempty"`
	Target              string    `json:"Target,omitempty"`
	TargetFaction       string    `json:"TargetFaction,omitempty"`
	TargetType          string    `json:"TargetType,omitempty"`
	TargetTypeLocalised string    `json:"TargetType_Localised,omitempty"`
	TargetLocalised     string    `json:"Target_Localised,omitempty"`
	Wing                bool      `json:"Wing,omitempty"`
	Event               string    `json:"event,omitempty"`
	Timestamp           time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMissionAccepted(ev *MissionAcceptedEvent) {
	return
}
