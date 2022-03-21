package file

import "time"

type BuyMicroResourcesEvent struct {
	Category      string    `json:"Category,omitempty"`
	Count         int       `json:"Count,omitempty"`
	MarketID      int       `json:"MarketID,omitempty"`
	Name          string    `json:"Name,omitempty"`
	NameLocalised string    `json:"Name_Localised,omitempty"`
	Price         int       `json:"Price,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBuyMicroResources(ev *BuyMicroResourcesEvent) {
	return
}
