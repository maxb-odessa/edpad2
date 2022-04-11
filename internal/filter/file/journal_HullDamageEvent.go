package file

import "time"

type HullDamageEvent struct {
	Fighter     bool      `mapstructure:"Fighter,omitempty"`
	Health      float64   `mapstructure:"Health,omitempty"`
	PlayerPilot bool      `mapstructure:"PlayerPilot,omitempty"`
	Event       string    `mapstructure:"event,omitempty"`
	Timestamp   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evHullDamage(ev *HullDamageEvent) {
	return
}
