package file

import "time"

type ModuleRetrieveEvent struct {
	EngineerModifications  string    `mapstructure:"EngineerModifications,omitempty"`
	Hot                    bool      `mapstructure:"Hot,omitempty"`
	Level                  int       `mapstructure:"Level,omitempty"`
	MarketID               int       `mapstructure:"MarketID,omitempty"`
	Quality                float64   `mapstructure:"Quality,omitempty"`
	RetrievedItem          string    `mapstructure:"RetrievedItem,omitempty"`
	RetrievedItemLocalised string    `mapstructure:"RetrievedItem_Localised,omitempty"`
	Ship                   string    `mapstructure:"Ship,omitempty"`
	ShipID                 int       `mapstructure:"ShipID,omitempty"`
	Slot                   string    `mapstructure:"Slot,omitempty"`
	SwapOutItem            string    `mapstructure:"SwapOutItem,omitempty"`
	SwapOutItemLocalised   string    `mapstructure:"SwapOutItem_Localised,omitempty"`
	Event                  string    `mapstructure:"event,omitempty"`
	Timestamp              time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evModuleRetrieve(ev *ModuleRetrieveEvent) {
	return
}
