package file

import "time"

type CommitCrimeEvent struct {
	Bounty    int       `mapstructure:"Bounty,omitempty"`
	CrimeType string    `mapstructure:"CrimeType,omitempty"`
	Faction   string    `mapstructure:"Faction,omitempty"`
	Fine      int       `mapstructure:"Fine,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCommitCrime(ev *CommitCrimeEvent) {
	return
}
