package file

import "time"

type MultiSellExplorationDataEvent struct {
	BaseValue  int `mapstructure:"BaseValue,omitempty"`
	Bonus      int `mapstructure:"Bonus,omitempty"`
	Discovered []struct {
		NumBodies  int    `mapstructure:"NumBodies,omitempty"`
		SystemName string `mapstructure:"SystemName,omitempty"`
	} `mapstructure:"Discovered,omitempty"`
	TotalEarnings int       `mapstructure:"TotalEarnings,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMultiSellExplorationData(ev *MultiSellExplorationDataEvent) {
	return
}
