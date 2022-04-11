package file

import "time"

type SellDronesEvent struct {
	Count     int       `mapstructure:"Count,omitempty"`
	SellPrice int       `mapstructure:"SellPrice,omitempty"`
	TotalSale int       `mapstructure:"TotalSale,omitempty"`
	Type      string    `mapstructure:"Type,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evSellDrones(ev *SellDronesEvent) {
	return
}
