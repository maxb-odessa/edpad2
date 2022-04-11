package file

import "time"

type BountyEvent struct {
	Reward  int `mapstructure:"Reward,omitempty"`
	Rewards []struct {
		Faction string `mapstructure:"Faction,omitempty"`
		Reward  int    `mapstructure:"Reward,omitempty"`
	} `mapstructure:"Rewards,omitempty"`
	Target                 string    `mapstructure:"Target,omitempty"`
	TargetLocalised        string    `mapstructure:"Target_Localised,omitempty"`
	TotalReward            int       `mapstructure:"TotalReward,omitempty"`
	VictimFaction          string    `mapstructure:"VictimFaction,omitempty"`
	VictimFactionLocalised string    `mapstructure:"VictimFaction_Localised,omitempty"`
	Event                  string    `mapstructure:"event,omitempty"`
	Timestamp              time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evBounty(ev *BountyEvent) {
	return
}
