package file

import "time"

type CommunityGoalRewardEvent struct {
	Cgid      int       `json:"CGID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Reward    int       `json:"Reward,omitempty"`
	System    string    `json:"System,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCommunityGoalReward(ev *CommunityGoalRewardEvent) {
	return
}
