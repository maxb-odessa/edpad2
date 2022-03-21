package file

import "time"

type RefuelPartialEvent struct {
	Amount    float64   `json:"Amount,omitempty"`
	Cost      int       `json:"Cost,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRefuelPartial(ev *RefuelPartialEvent) {
	return
}
