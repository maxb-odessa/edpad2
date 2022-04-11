package file

import "time"

type DiedEvent struct {
	KillerName string    `mapstructure:"KillerName,omitempty"`
	KillerRank string    `mapstructure:"KillerRank,omitempty"`
	KillerShip string    `mapstructure:"KillerShip,omitempty"`
	Event      string    `mapstructure:"event,omitempty"`
	Timestamp  time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evDied(ev *DiedEvent) {
	return
}
