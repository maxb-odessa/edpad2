package file

import "time"

type RefuelAllEvent struct {
	Amount    float64   `mapstructure:"Amount,omitempty"`
	Cost      int       `mapstructure:"Cost,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evRefuelAll(ev *RefuelAllEvent) {
	return
}
