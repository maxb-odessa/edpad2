package file

import "time"

type EngineerContributionEvent struct {
	Commodity          string    `mapstructure:"Commodity,omitempty"`
	CommodityLocalised string    `mapstructure:"Commodity_Localised,omitempty"`
	Engineer           string    `mapstructure:"Engineer,omitempty"`
	EngineerID         int       `mapstructure:"EngineerID,omitempty"`
	Quantity           int       `mapstructure:"Quantity,omitempty"`
	TotalQuantity      int       `mapstructure:"TotalQuantity,omitempty"`
	Type               string    `mapstructure:"Type,omitempty"`
	Event              string    `mapstructure:"event,omitempty"`
	Timestamp          time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evEngineerContribution(ev *EngineerContributionEvent) {
	return
}
