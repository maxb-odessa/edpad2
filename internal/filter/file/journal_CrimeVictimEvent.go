package file

import "time"

type CrimeVictimEvent struct {
	Bounty    int       `json:"Bounty,omitempty"`
	CrimeType string    `json:"CrimeType,omitempty"`
	Fine      int       `json:"Fine,omitempty"`
	Offender  string    `json:"Offender,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCrimeVictim(ev *CrimeVictimEvent) {
	return
}
