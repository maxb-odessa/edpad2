package file

import "time"

type InterdictedEvent struct {
	CombatRank  int       `json:"CombatRank,omitempty"`
	Faction     string    `json:"Faction,omitempty"`
	Interdictor string    `json:"Interdictor,omitempty"`
	IsPlayer    bool      `json:"IsPlayer,omitempty"`
	Submitted   bool      `json:"Submitted,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evInterdicted(ev *InterdictedEvent) {
	return
}
