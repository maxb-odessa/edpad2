package file

import "time"

type SellSuitEvent struct {
	Name          string        `mapstructure:"Name,omitempty"`
	NameLocalised string        `mapstructure:"Name_Localised,omitempty"`
	Price         int           `mapstructure:"Price,omitempty"`
	SuitID        int           `mapstructure:"SuitID,omitempty"`
	SuitMods      []interface{} `mapstructure:"SuitMods,omitempty"`
	Event         string        `mapstructure:"event,omitempty"`
	Timestamp     time.Time     `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evSellSuit(ev *SellSuitEvent) {
	return
}
