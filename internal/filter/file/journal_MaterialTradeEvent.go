package file

import "time"

type MaterialTradeEvent struct {
	MarketID int `mapstructure:"MarketID,omitempty"`
	Paid     struct {
		Category          string `mapstructure:"Category,omitempty"`
		Material          string `mapstructure:"Material,omitempty"`
		MaterialLocalised string `mapstructure:"Material_Localised,omitempty"`
		Quantity          int    `mapstructure:"Quantity,omitempty"`
	} `mapstructure:"Paid,omitempty"`
	Received struct {
		Category          string `mapstructure:"Category,omitempty"`
		Material          string `mapstructure:"Material,omitempty"`
		MaterialLocalised string `mapstructure:"Material_Localised,omitempty"`
		Quantity          int    `mapstructure:"Quantity,omitempty"`
	} `mapstructure:"Received,omitempty"`
	TraderType string    `mapstructure:"TraderType,omitempty"`
	Event      string    `mapstructure:"event,omitempty"`
	Timestamp  time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMaterialTrade(ev *MaterialTradeEvent) {
	return
}
