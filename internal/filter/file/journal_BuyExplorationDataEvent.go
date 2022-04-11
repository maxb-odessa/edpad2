package file

import "time"

type BuyExplorationDataEvent struct {
	Cost      int       `mapstructure:"Cost,omitempty"`
	System    string    `mapstructure:"System,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evBuyExplorationData(ev *BuyExplorationDataEvent) {
	return
}
