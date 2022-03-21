package file

import "time"

type BuyTradeDataEvent struct {
	Cost      int       `json:"Cost,omitempty"`
	System    string    `json:"System,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBuyTradeData(ev *BuyTradeDataEvent) {
	return
}
