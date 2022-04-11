package file

import "time"

type BuyDronesEvent struct {
	BuyPrice  int       `mapstructure:"BuyPrice,omitempty"`
	Count     int       `mapstructure:"Count,omitempty"`
	TotalCost int       `mapstructure:"TotalCost,omitempty"`
	Type      string    `mapstructure:"Type,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evBuyDrones(ev *BuyDronesEvent) {
	return
}
