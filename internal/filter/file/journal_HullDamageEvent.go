package file

import "time"

type HullDamageEvent struct {
	Fighter     bool      `json:"Fighter,omitempty"`
	Health      float64   `json:"Health,omitempty"`
	PlayerPilot bool      `json:"PlayerPilot,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evHullDamage(ev *HullDamageEvent) {
	return
}
