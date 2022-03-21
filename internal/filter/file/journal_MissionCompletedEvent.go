package file

import "time"

type MissionCompletedEvent struct {
	Commodity       string `json:"Commodity,omitempty"`
	CommodityReward []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"CommodityReward,omitempty"`
	CommodityLocalised string `json:"Commodity_Localised,omitempty"`
	Count              int    `json:"Count,omitempty"`
	DestinationStation string `json:"DestinationStation,omitempty"`
	DestinationSystem  string `json:"DestinationSystem,omitempty"`
	Donated            int    `json:"Donated,omitempty"`
	Donation           string `json:"Donation,omitempty"`
	Faction            string `json:"Faction,omitempty"`
	FactionEffects     []struct {
		Effects []struct {
			Effect          string `json:"Effect,omitempty"`
			EffectLocalised string `json:"Effect_Localised,omitempty"`
			Trend           string `json:"Trend,omitempty"`
		} `json:"Effects,omitempty"`
		Faction   string `json:"Faction,omitempty"`
		Influence []struct {
			Influence     string `json:"Influence,omitempty"`
			SystemAddress int    `json:"SystemAddress,omitempty"`
			Trend         string `json:"Trend,omitempty"`
		} `json:"Influence,omitempty"`
		Reputation      string `json:"Reputation,omitempty"`
		ReputationTrend string `json:"ReputationTrend,omitempty"`
	} `json:"FactionEffects,omitempty"`
	MaterialsReward []struct {
		Category          string `json:"Category,omitempty"`
		CategoryLocalised string `json:"Category_Localised,omitempty"`
		Count             int    `json:"Count,omitempty"`
		Name              string `json:"Name,omitempty"`
		NameLocalised     string `json:"Name_Localised,omitempty"`
	} `json:"MaterialsReward,omitempty"`
	MissionID            int       `json:"MissionID,omitempty"`
	Name                 string    `json:"Name,omitempty"`
	NewDestinationSystem string    `json:"NewDestinationSystem,omitempty"`
	Reward               int       `json:"Reward,omitempty"`
	Target               string    `json:"Target,omitempty"`
	TargetFaction        string    `json:"TargetFaction,omitempty"`
	TargetType           string    `json:"TargetType,omitempty"`
	TargetTypeLocalised  string    `json:"TargetType_Localised,omitempty"`
	TargetLocalised      string    `json:"Target_Localised,omitempty"`
	Event                string    `json:"event,omitempty"`
	Timestamp            time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMissionCompleted(ev *MissionCompletedEvent) {
	return
}
