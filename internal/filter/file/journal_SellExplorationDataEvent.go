package file

import "time"

type SellExplorationDataEvent struct {
	BaseValue     int           `mapstructure:"BaseValue,omitempty"`
	Bonus         int           `mapstructure:"Bonus,omitempty"`
	Discovered    []interface{} `mapstructure:"Discovered,omitempty"`
	Systems       []string      `mapstructure:"Systems,omitempty"`
	TotalEarnings int           `mapstructure:"TotalEarnings,omitempty"`
	Event         string        `mapstructure:"event,omitempty"`
	Timestamp     time.Time     `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evSellExplorationData(ev *SellExplorationDataEvent) {
	return
}
