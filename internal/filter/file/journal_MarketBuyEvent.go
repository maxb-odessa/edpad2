package file

import "time"

type MarketBuyEvent struct {
	BuyPrice      int       `mapstructure:"BuyPrice,omitempty"`
	Count         int       `mapstructure:"Count,omitempty"`
	MarketID      int       `mapstructure:"MarketID,omitempty"`
	TotalCost     int       `mapstructure:"TotalCost,omitempty"`
	Type          string    `mapstructure:"Type,omitempty"`
	TypeLocalised string    `mapstructure:"Type_Localised,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMarketBuy(ev *MarketBuyEvent) {
	return
}
