package file

import "time"

type CommitCrimeEvent struct {
	Bounty    int       `json:"Bounty,omitempty"`
	CrimeType string    `json:"CrimeType,omitempty"`
	Faction   string    `json:"Faction,omitempty"`
	Fine      int       `json:"Fine,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCommitCrime(ev *CommitCrimeEvent) {
	return
}
