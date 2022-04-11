package file

import "time"

type ModuleStoreEvent struct {
	EngineerModifications string    `mapstructure:"EngineerModifications,omitempty"`
	Hot                   bool      `mapstructure:"Hot,omitempty"`
	Level                 int       `mapstructure:"Level,omitempty"`
	MarketID              int       `mapstructure:"MarketID,omitempty"`
	Quality               float64   `mapstructure:"Quality,omitempty"`
	Ship                  string    `mapstructure:"Ship,omitempty"`
	ShipID                int       `mapstructure:"ShipID,omitempty"`
	Slot                  string    `mapstructure:"Slot,omitempty"`
	StoredItem            string    `mapstructure:"StoredItem,omitempty"`
	StoredItemLocalised   string    `mapstructure:"StoredItem_Localised,omitempty"`
	Event                 string    `mapstructure:"event,omitempty"`
	Timestamp             time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evModuleStore(ev *ModuleStoreEvent) {
	return
}
