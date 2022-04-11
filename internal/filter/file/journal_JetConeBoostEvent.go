package file

import "time"

type JetConeBoostEvent struct {
	BoostValue float64   `mapstructure:"BoostValue,omitempty"`
	Event      string    `mapstructure:"event,omitempty"`
	Timestamp  time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evJetConeBoost(ev *JetConeBoostEvent) {
	return
}
