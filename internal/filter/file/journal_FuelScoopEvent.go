package file

import "time"

type FuelScoopEvent struct {
	Scooped   float64   `json:"Scooped,omitempty"`
	Total     float64   `json:"Total,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFuelScoop(ev *FuelScoopEvent) {
	return
}
