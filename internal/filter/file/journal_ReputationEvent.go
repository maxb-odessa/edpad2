package file

import "time"

type ReputationEvent struct {
	Alliance   float64   `json:"Alliance,omitempty"`
	Empire     float64   `json:"Empire,omitempty"`
	Federation float64   `json:"Federation,omitempty"`
	Event      string    `json:"event,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evReputation(ev *ReputationEvent) {
	return
}
