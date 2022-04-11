package file

import "time"

type ModuleSwapEvent struct {
	FromItem          string    `mapstructure:"FromItem,omitempty"`
	FromItemLocalised string    `mapstructure:"FromItem_Localised,omitempty"`
	FromSlot          string    `mapstructure:"FromSlot,omitempty"`
	MarketID          int       `mapstructure:"MarketID,omitempty"`
	Ship              string    `mapstructure:"Ship,omitempty"`
	ShipID            int       `mapstructure:"ShipID,omitempty"`
	ToItem            string    `mapstructure:"ToItem,omitempty"`
	ToSlot            string    `mapstructure:"ToSlot,omitempty"`
	Event             string    `mapstructure:"event,omitempty"`
	Timestamp         time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evModuleSwap(ev *ModuleSwapEvent) {
	return
}
