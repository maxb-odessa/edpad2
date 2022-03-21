package file

import "time"

type InvitedToSquadronEvent struct {
	SquadronName string    `json:"SquadronName,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evInvitedToSquadron(ev *InvitedToSquadronEvent) {
	return
}
