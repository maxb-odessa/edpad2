package file

import "time"

type LaunchFighterEvent struct {
	ID               int       `mapstructure:"ID,omitempty"`
	Loadout          string    `mapstructure:"Loadout,omitempty"`
	PlayerControlled bool      `mapstructure:"PlayerControlled,omitempty"`
	Event            string    `mapstructure:"event,omitempty"`
	Timestamp        time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evLaunchFighter(ev *LaunchFighterEvent) {
	return
}
