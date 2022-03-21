package file

import "time"

type ModuleSwapEvent struct {
	FromItem          string    `json:"FromItem,omitempty"`
	FromItemLocalised string    `json:"FromItem_Localised,omitempty"`
	FromSlot          string    `json:"FromSlot,omitempty"`
	MarketID          int       `json:"MarketID,omitempty"`
	Ship              string    `json:"Ship,omitempty"`
	ShipID            int       `json:"ShipID,omitempty"`
	ToItem            string    `json:"ToItem,omitempty"`
	ToSlot            string    `json:"ToSlot,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evModuleSwap(ev *ModuleSwapEvent) {
	return
}
