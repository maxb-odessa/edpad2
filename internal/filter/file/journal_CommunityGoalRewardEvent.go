package file

import "time"

type CommunityGoalRewardEvent struct {
	Cgid      int       `mapstructure:"CGID,omitempty"`
	Name      string    `mapstructure:"Name,omitempty"`
	Reward    int       `mapstructure:"Reward,omitempty"`
	System    string    `mapstructure:"System,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCommunityGoalReward(ev *CommunityGoalRewardEvent) {
	return
}
