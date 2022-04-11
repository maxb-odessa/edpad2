package file

import "time"

type ReputationEvent struct {
	Alliance   float64   `mapstructure:"Alliance,omitempty"`
	Empire     float64   `mapstructure:"Empire,omitempty"`
	Federation float64   `mapstructure:"Federation,omitempty"`
	Event      string    `mapstructure:"event,omitempty"`
	Timestamp  time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evReputation(ev *ReputationEvent) {
	return
}
