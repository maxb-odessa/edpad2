package file

import "time"

type MultiSellExplorationDataEvent struct {
	BaseValue  int `json:"BaseValue,omitempty"`
	Bonus      int `json:"Bonus,omitempty"`
	Discovered []struct {
		NumBodies  int    `json:"NumBodies,omitempty"`
		SystemName string `json:"SystemName,omitempty"`
	} `json:"Discovered,omitempty"`
	TotalEarnings int       `json:"TotalEarnings,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMultiSellExplorationData(ev *MultiSellExplorationDataEvent) {
	return
}
