package file

import "time"

type SellSuitEvent struct {
	Name          string        `json:"Name,omitempty"`
	NameLocalised string        `json:"Name_Localised,omitempty"`
	Price         int           `json:"Price,omitempty"`
	SuitID        int           `json:"SuitID,omitempty"`
	SuitMods      []interface{} `json:"SuitMods,omitempty"`
	Event         string        `json:"event,omitempty"`
	Timestamp     time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evSellSuit(ev *SellSuitEvent) {
	return
}
