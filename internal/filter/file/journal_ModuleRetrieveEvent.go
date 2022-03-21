package file

import "time"

type ModuleRetrieveEvent struct {
	EngineerModifications  string    `json:"EngineerModifications,omitempty"`
	Hot                    bool      `json:"Hot,omitempty"`
	Level                  int       `json:"Level,omitempty"`
	MarketID               int       `json:"MarketID,omitempty"`
	Quality                float64   `json:"Quality,omitempty"`
	RetrievedItem          string    `json:"RetrievedItem,omitempty"`
	RetrievedItemLocalised string    `json:"RetrievedItem_Localised,omitempty"`
	Ship                   string    `json:"Ship,omitempty"`
	ShipID                 int       `json:"ShipID,omitempty"`
	Slot                   string    `json:"Slot,omitempty"`
	SwapOutItem            string    `json:"SwapOutItem,omitempty"`
	SwapOutItemLocalised   string    `json:"SwapOutItem_Localised,omitempty"`
	Event                  string    `json:"event,omitempty"`
	Timestamp              time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evModuleRetrieve(ev *ModuleRetrieveEvent) {
	return
}
