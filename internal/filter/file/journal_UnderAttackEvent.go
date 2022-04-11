package file

import "time"

type UnderAttackEvent struct {
	Target    string    `mapstructure:"Target,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evUnderAttack(ev *UnderAttackEvent) {
	return
}
