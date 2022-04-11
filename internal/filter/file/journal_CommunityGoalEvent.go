package file

import "time"

type CommunityGoalEvent struct {
	CurrentGoals []struct {
		Bonus                int       `mapstructure:"Bonus,omitempty"`
		Cgid                 int       `mapstructure:"CGID,omitempty"`
		CurrentTotal         int       `mapstructure:"CurrentTotal,omitempty"`
		Expiry               time.Time `mapstructure:"Expiry,omitempty"`
		IsComplete           bool      `mapstructure:"IsComplete,omitempty"`
		MarketName           string    `mapstructure:"MarketName,omitempty"`
		NumContributors      int       `mapstructure:"NumContributors,omitempty"`
		PlayerContribution   int       `mapstructure:"PlayerContribution,omitempty"`
		PlayerInTopRank      bool      `mapstructure:"PlayerInTopRank,omitempty"`
		PlayerPercentileBand int       `mapstructure:"PlayerPercentileBand,omitempty"`
		SystemName           string    `mapstructure:"SystemName,omitempty"`
		TierReached          string    `mapstructure:"TierReached,omitempty"`
		Title                string    `mapstructure:"Title,omitempty"`
		TopRankSize          int       `mapstructure:"TopRankSize,omitempty"`
		TopTier              struct {
			Bonus string `mapstructure:"Bonus,omitempty"`
			Name  string `mapstructure:"Name,omitempty"`
		} `mapstructure:"TopTier,omitempty"`
	} `mapstructure:"CurrentGoals,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCommunityGoal(ev *CommunityGoalEvent) {
	return
}
