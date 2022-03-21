package file

import "time"

type DiedEvent struct {
	KillerName string    `json:"KillerName,omitempty"`
	KillerRank string    `json:"KillerRank,omitempty"`
	KillerShip string    `json:"KillerShip,omitempty"`
	Event      string    `json:"event,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDied(ev *DiedEvent) {
	return
}
