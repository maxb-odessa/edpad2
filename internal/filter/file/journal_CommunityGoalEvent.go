package file

import "time"

type CommunityGoalEvent struct {
	CurrentGoals []struct {
		Bonus                int       `json:"Bonus,omitempty"`
		Cgid                 int       `json:"CGID,omitempty"`
		CurrentTotal         int       `json:"CurrentTotal,omitempty"`
		Expiry               time.Time `json:"Expiry,omitempty"`
		IsComplete           bool      `json:"IsComplete,omitempty"`
		MarketName           string    `json:"MarketName,omitempty"`
		NumContributors      int       `json:"NumContributors,omitempty"`
		PlayerContribution   int       `json:"PlayerContribution,omitempty"`
		PlayerInTopRank      bool      `json:"PlayerInTopRank,omitempty"`
		PlayerPercentileBand int       `json:"PlayerPercentileBand,omitempty"`
		SystemName           string    `json:"SystemName,omitempty"`
		TierReached          string    `json:"TierReached,omitempty"`
		Title                string    `json:"Title,omitempty"`
		TopRankSize          int       `json:"TopRankSize,omitempty"`
		TopTier              struct {
			Bonus string `json:"Bonus,omitempty"`
			Name  string `json:"Name,omitempty"`
		} `json:"TopTier,omitempty"`
	} `json:"CurrentGoals,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCommunityGoal(ev *CommunityGoalEvent) {
	return
}
