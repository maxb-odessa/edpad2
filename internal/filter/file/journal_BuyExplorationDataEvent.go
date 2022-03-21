package file

import "time"

type BuyExplorationDataEvent struct {
	Cost      int       `json:"Cost,omitempty"`
	System    string    `json:"System,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBuyExplorationData(ev *BuyExplorationDataEvent) {
	return
}
