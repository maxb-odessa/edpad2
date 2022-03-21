package file

import "time"

type MarketSellEvent struct {
	AvgPricePaid  int       `json:"AvgPricePaid,omitempty"`
	BlackMarket   bool      `json:"BlackMarket,omitempty"`
	Count         int       `json:"Count,omitempty"`
	IllegalGoods  bool      `json:"IllegalGoods,omitempty"`
	MarketID      int       `json:"MarketID,omitempty"`
	SellPrice     int       `json:"SellPrice,omitempty"`
	StolenGoods   bool      `json:"StolenGoods,omitempty"`
	TotalSale     int       `json:"TotalSale,omitempty"`
	Type          string    `json:"Type,omitempty"`
	TypeLocalised string    `json:"Type_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMarketSell(ev *MarketSellEvent) {
	return
}
