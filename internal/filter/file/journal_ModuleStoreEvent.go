package file

import "time"

type ModuleStoreEvent struct {
	EngineerModifications string    `json:"EngineerModifications,omitempty"`
	Hot                   bool      `json:"Hot,omitempty"`
	Level                 int       `json:"Level,omitempty"`
	MarketID              int       `json:"MarketID,omitempty"`
	Quality               float64   `json:"Quality,omitempty"`
	Ship                  string    `json:"Ship,omitempty"`
	ShipID                int       `json:"ShipID,omitempty"`
	Slot                  string    `json:"Slot,omitempty"`
	StoredItem            string    `json:"StoredItem,omitempty"`
	StoredItemLocalised   string    `json:"StoredItem_Localised,omitempty"`
	Event                 string    `json:"event,omitempty"`
	Timestamp             time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evModuleStore(ev *ModuleStoreEvent) {
	return
}
