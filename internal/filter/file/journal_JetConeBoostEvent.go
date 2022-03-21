package file

import "time"

type JetConeBoostEvent struct {
	BoostValue float64   `json:"BoostValue,omitempty"`
	Event      string    `json:"event,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evJetConeBoost(ev *JetConeBoostEvent) {
	return
}
