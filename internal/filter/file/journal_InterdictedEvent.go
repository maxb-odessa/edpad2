package file

import "time"

type InterdictedEvent struct {
	CombatRank  int       `mapstructure:"CombatRank,omitempty"`
	Faction     string    `mapstructure:"Faction,omitempty"`
	Interdictor string    `mapstructure:"Interdictor,omitempty"`
	IsPlayer    bool      `mapstructure:"IsPlayer,omitempty"`
	Submitted   bool      `mapstructure:"Submitted,omitempty"`
	Event       string    `mapstructure:"event,omitempty"`
	Timestamp   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evInterdicted(ev *InterdictedEvent) {
	return
}
