package file

import "time"

type MarketSellEvent struct {
	AvgPricePaid  int       `mapstructure:"AvgPricePaid,omitempty"`
	BlackMarket   bool      `mapstructure:"BlackMarket,omitempty"`
	Count         int       `mapstructure:"Count,omitempty"`
	IllegalGoods  bool      `mapstructure:"IllegalGoods,omitempty"`
	MarketID      int       `mapstructure:"MarketID,omitempty"`
	SellPrice     int       `mapstructure:"SellPrice,omitempty"`
	StolenGoods   bool      `mapstructure:"StolenGoods,omitempty"`
	TotalSale     int       `mapstructure:"TotalSale,omitempty"`
	Type          string    `mapstructure:"Type,omitempty"`
	TypeLocalised string    `mapstructure:"Type_Localised,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMarketSell(ev *MarketSellEvent) {
	return
}
