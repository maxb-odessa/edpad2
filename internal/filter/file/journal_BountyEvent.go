package file

import "time"

type BountyEvent struct {
	Reward  int `json:"Reward,omitempty"`
	Rewards []struct {
		Faction string `json:"Faction,omitempty"`
		Reward  int    `json:"Reward,omitempty"`
	} `json:"Rewards,omitempty"`
	Target                 string    `json:"Target,omitempty"`
	TargetLocalised        string    `json:"Target_Localised,omitempty"`
	TotalReward            int       `json:"TotalReward,omitempty"`
	VictimFaction          string    `json:"VictimFaction,omitempty"`
	VictimFactionLocalised string    `json:"VictimFaction_Localised,omitempty"`
	Event                  string    `json:"event,omitempty"`
	Timestamp              time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBounty(ev *BountyEvent) {
	return
}
