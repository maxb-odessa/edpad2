package file

import "time"

type RefuelAllEvent struct {
	Amount    float64   `json:"Amount,omitempty"`
	Cost      int       `json:"Cost,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRefuelAll(ev *RefuelAllEvent) {
	return
}
