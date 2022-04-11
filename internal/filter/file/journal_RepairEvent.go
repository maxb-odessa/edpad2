package file

import "time"

type RepairEvent struct {
	Cost      int       `mapstructure:"Cost,omitempty"`
	Item      string    `mapstructure:"Item,omitempty"`
	Items     []string  `mapstructure:"Items,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evRepair(ev *RepairEvent) {
	return
}
