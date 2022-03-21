package file

import "time"

type SellDronesEvent struct {
	Count     int       `json:"Count,omitempty"`
	SellPrice int       `json:"SellPrice,omitempty"`
	TotalSale int       `json:"TotalSale,omitempty"`
	Type      string    `json:"Type,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSellDrones(ev *SellDronesEvent) {
	return
}
