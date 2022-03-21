package file

import "time"

type UnderAttackEvent struct {
	Target    string    `json:"Target,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evUnderAttack(ev *UnderAttackEvent) {
	return
}
