package file

import "time"

type RepairAllEvent struct {
	Cost      int       `json:"Cost,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRepairAll(ev *RepairAllEvent) {
	return
}
