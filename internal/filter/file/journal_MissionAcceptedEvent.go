package file

import "time"

type MissionAcceptedEvent struct {
	Commodity           string    `mapstructure:"Commodity,omitempty"`
	CommodityLocalised  string    `mapstructure:"Commodity_Localised,omitempty"`
	Count               int       `mapstructure:"Count,omitempty"`
	DestinationStation  string    `mapstructure:"DestinationStation,omitempty"`
	DestinationSystem   string    `mapstructure:"DestinationSystem,omitempty"`
	Donation            string    `mapstructure:"Donation,omitempty"`
	Expiry              time.Time `mapstructure:"Expiry,omitempty"`
	Faction             string    `mapstructure:"Faction,omitempty"`
	Influence           string    `mapstructure:"Influence,omitempty"`
	LocalisedName       string    `mapstructure:"LocalisedName,omitempty"`
	MissionID           int       `mapstructure:"MissionID,omitempty"`
	Name                string    `mapstructure:"Name,omitempty"`
	PassengerCount      int       `mapstructure:"PassengerCount,omitempty"`
	PassengerType       string    `mapstructure:"PassengerType,omitempty"`
	PassengerViPs       bool      `mapstructure:"PassengerVIPs,omitempty"`
	PassengerWanted     bool      `mapstructure:"PassengerWanted,omitempty"`
	Reputation          string    `mapstructure:"Reputation,omitempty"`
	Reward              int       `mapstructure:"Reward,omitempty"`
	Target              string    `mapstructure:"Target,omitempty"`
	TargetFaction       string    `mapstructure:"TargetFaction,omitempty"`
	TargetType          string    `mapstructure:"TargetType,omitempty"`
	TargetTypeLocalised string    `mapstructure:"TargetType_Localised,omitempty"`
	TargetLocalised     string    `mapstructure:"Target_Localised,omitempty"`
	Wing                bool      `mapstructure:"Wing,omitempty"`
	Event               string    `mapstructure:"event,omitempty"`
	Timestamp           time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMissionAccepted(ev *MissionAcceptedEvent) {
	return
}
