package file

import "time"

type RepairAllEvent struct {
	Cost      int       `mapstructure:"Cost,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evRepairAll(ev *RepairAllEvent) {
	return
}
