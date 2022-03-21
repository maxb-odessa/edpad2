package file

import "time"

type RepairEvent struct {
	Cost      int       `json:"Cost,omitempty"`
	Item      string    `json:"Item,omitempty"`
	Items     []string  `json:"Items,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRepair(ev *RepairEvent) {
	return
}
