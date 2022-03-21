package file

import "time"

type CrewHireEvent struct {
	CombatRank int       `json:"CombatRank,omitempty"`
	Cost       int       `json:"Cost,omitempty"`
	CrewID     int       `json:"CrewID,omitempty"`
	Faction    string    `json:"Faction,omitempty"`
	Name       string    `json:"Name,omitempty"`
	Event      string    `json:"event,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCrewHire(ev *CrewHireEvent) {
	return
}
