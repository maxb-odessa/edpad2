package file

import "time"

type SquadronStartupEvent struct {
	CurrentRank  int       `json:"CurrentRank,omitempty"`
	SquadronName string    `json:"SquadronName,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSquadronStartup(ev *SquadronStartupEvent) {
	return
}
