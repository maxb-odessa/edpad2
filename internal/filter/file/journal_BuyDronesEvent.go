package file

import "time"

type BuyDronesEvent struct {
	BuyPrice  int       `json:"BuyPrice,omitempty"`
	Count     int       `json:"Count,omitempty"`
	TotalCost int       `json:"TotalCost,omitempty"`
	Type      string    `json:"Type,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBuyDrones(ev *BuyDronesEvent) {
	return
}
