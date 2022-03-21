package file

import "time"

type EngineerContributionEvent struct {
	Commodity          string    `json:"Commodity,omitempty"`
	CommodityLocalised string    `json:"Commodity_Localised,omitempty"`
	Engineer           string    `json:"Engineer,omitempty"`
	EngineerID         int       `json:"EngineerID,omitempty"`
	Quantity           int       `json:"Quantity,omitempty"`
	TotalQuantity      int       `json:"TotalQuantity,omitempty"`
	Type               string    `json:"Type,omitempty"`
	Event              string    `json:"event,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evEngineerContribution(ev *EngineerContributionEvent) {
	return
}
