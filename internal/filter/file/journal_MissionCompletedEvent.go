package file

import "time"

type MissionCompletedEvent struct {
	Commodity       string `mapstructure:"Commodity,omitempty"`
	CommodityReward []struct {
		Count         int    `mapstructure:"Count,omitempty"`
		Name          string `mapstructure:"Name,omitempty"`
		NameLocalised string `mapstructure:"Name_Localised,omitempty"`
	} `mapstructure:"CommodityReward,omitempty"`
	CommodityLocalised string `mapstructure:"Commodity_Localised,omitempty"`
	Count              int    `mapstructure:"Count,omitempty"`
	DestinationStation string `mapstructure:"DestinationStation,omitempty"`
	DestinationSystem  string `mapstructure:"DestinationSystem,omitempty"`
	Donated            int    `mapstructure:"Donated,omitempty"`
	Donation           string `mapstructure:"Donation,omitempty"`
	Faction            string `mapstructure:"Faction,omitempty"`
	FactionEffects     []struct {
		Effects []struct {
			Effect          string `mapstructure:"Effect,omitempty"`
			EffectLocalised string `mapstructure:"Effect_Localised,omitempty"`
			Trend           string `mapstructure:"Trend,omitempty"`
		} `mapstructure:"Effects,omitempty"`
		Faction   string `mapstructure:"Faction,omitempty"`
		Influence []struct {
			Influence     string `mapstructure:"Influence,omitempty"`
			SystemAddress int    `mapstructure:"SystemAddress,omitempty"`
			Trend         string `mapstructure:"Trend,omitempty"`
		} `mapstructure:"Influence,omitempty"`
		Reputation      string `mapstructure:"Reputation,omitempty"`
		ReputationTrend string `mapstructure:"ReputationTrend,omitempty"`
	} `mapstructure:"FactionEffects,omitempty"`
	MaterialsReward []struct {
		Category          string `mapstructure:"Category,omitempty"`
		CategoryLocalised string `mapstructure:"Category_Localised,omitempty"`
		Count             int    `mapstructure:"Count,omitempty"`
		Name              string `mapstructure:"Name,omitempty"`
		NameLocalised     string `mapstructure:"Name_Localised,omitempty"`
	} `mapstructure:"MaterialsReward,omitempty"`
	MissionID            int       `mapstructure:"MissionID,omitempty"`
	Name                 string    `mapstructure:"Name,omitempty"`
	NewDestinationSystem string    `mapstructure:"NewDestinationSystem,omitempty"`
	Reward               int       `mapstructure:"Reward,omitempty"`
	Target               string    `mapstructure:"Target,omitempty"`
	TargetFaction        string    `mapstructure:"TargetFaction,omitempty"`
	TargetType           string    `mapstructure:"TargetType,omitempty"`
	TargetTypeLocalised  string    `mapstructure:"TargetType_Localised,omitempty"`
	TargetLocalised      string    `mapstructure:"Target_Localised,omitempty"`
	Event                string    `mapstructure:"event,omitempty"`
	Timestamp            time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMissionCompleted(ev *MissionCompletedEvent) {
	return
}
