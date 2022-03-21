package file

import "time"

type MarketBuyEvent struct {
	BuyPrice      int       `json:"BuyPrice,omitempty"`
	Count         int       `json:"Count,omitempty"`
	MarketID      int       `json:"MarketID,omitempty"`
	TotalCost     int       `json:"TotalCost,omitempty"`
	Type          string    `json:"Type,omitempty"`
	TypeLocalised string    `json:"Type_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMarketBuy(ev *MarketBuyEvent) {
	return
}
