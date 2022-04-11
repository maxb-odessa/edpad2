package file

import "time"

type InvitedToSquadronEvent struct {
	SquadronName string    `mapstructure:"SquadronName,omitempty"`
	Event        string    `mapstructure:"event,omitempty"`
	Timestamp    time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evInvitedToSquadron(ev *InvitedToSquadronEvent) {
	return
}
