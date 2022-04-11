package file

import "time"

type RefuelPartialEvent struct {
	Amount    float64   `mapstructure:"Amount,omitempty"`
	Cost      int       `mapstructure:"Cost,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evRefuelPartial(ev *RefuelPartialEvent) {
	return
}
