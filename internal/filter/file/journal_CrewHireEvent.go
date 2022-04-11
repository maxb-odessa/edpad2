package file

import "time"

type CrewHireEvent struct {
	CombatRank int       `mapstructure:"CombatRank,omitempty"`
	Cost       int       `mapstructure:"Cost,omitempty"`
	CrewID     int       `mapstructure:"CrewID,omitempty"`
	Faction    string    `mapstructure:"Faction,omitempty"`
	Name       string    `mapstructure:"Name,omitempty"`
	Event      string    `mapstructure:"event,omitempty"`
	Timestamp  time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCrewHire(ev *CrewHireEvent) {
	return
}
