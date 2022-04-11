package file

import "time"

type SquadronStartupEvent struct {
	CurrentRank  int       `mapstructure:"CurrentRank,omitempty"`
	SquadronName string    `mapstructure:"SquadronName,omitempty"`
	Event        string    `mapstructure:"event,omitempty"`
	Timestamp    time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evSquadronStartup(ev *SquadronStartupEvent) {
	return
}
