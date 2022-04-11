package file

import "time"

type BuyMicroResourcesEvent struct {
	Category      string    `mapstructure:"Category,omitempty"`
	Count         int       `mapstructure:"Count,omitempty"`
	MarketID      int       `mapstructure:"MarketID,omitempty"`
	Name          string    `mapstructure:"Name,omitempty"`
	NameLocalised string    `mapstructure:"Name_Localised,omitempty"`
	Price         int       `mapstructure:"Price,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evBuyMicroResources(ev *BuyMicroResourcesEvent) {
	return
}
