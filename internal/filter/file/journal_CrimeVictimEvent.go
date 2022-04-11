package file

import "time"

type CrimeVictimEvent struct {
	Bounty    int       `mapstructure:"Bounty,omitempty"`
	CrimeType string    `mapstructure:"CrimeType,omitempty"`
	Fine      int       `mapstructure:"Fine,omitempty"`
	Offender  string    `mapstructure:"Offender,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCrimeVictim(ev *CrimeVictimEvent) {
	return
}
