package file

import "time"

type MaterialTradeEvent struct {
	MarketID int `json:"MarketID,omitempty"`
	Paid     struct {
		Category          string `json:"Category,omitempty"`
		Material          string `json:"Material,omitempty"`
		MaterialLocalised string `json:"Material_Localised,omitempty"`
		Quantity          int    `json:"Quantity,omitempty"`
	} `json:"Paid,omitempty"`
	Received struct {
		Category          string `json:"Category,omitempty"`
		Material          string `json:"Material,omitempty"`
		MaterialLocalised string `json:"Material_Localised,omitempty"`
		Quantity          int    `json:"Quantity,omitempty"`
	} `json:"Received,omitempty"`
	TraderType string    `json:"TraderType,omitempty"`
	Event      string    `json:"event,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMaterialTrade(ev *MaterialTradeEvent) {
	return
}
