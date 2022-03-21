package file

import "time"

type SellExplorationDataEvent struct {
	BaseValue     int           `json:"BaseValue,omitempty"`
	Bonus         int           `json:"Bonus,omitempty"`
	Discovered    []interface{} `json:"Discovered,omitempty"`
	Systems       []string      `json:"Systems,omitempty"`
	TotalEarnings int           `json:"TotalEarnings,omitempty"`
	Event         string        `json:"event,omitempty"`
	Timestamp     time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evSellExplorationData(ev *SellExplorationDataEvent) {
	return
}
