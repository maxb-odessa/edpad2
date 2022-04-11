package file

import "time"

type StoredModulesEvent struct {
	Items []struct {
		BuyPrice              int     `mapstructure:"BuyPrice,omitempty"`
		EngineerModifications string  `mapstructure:"EngineerModifications,omitempty"`
		Hot                   bool    `mapstructure:"Hot,omitempty"`
		InTransit             bool    `mapstructure:"InTransit,omitempty"`
		Level                 int     `mapstructure:"Level,omitempty"`
		MarketID              int     `mapstructure:"MarketID,omitempty"`
		Name                  string  `mapstructure:"Name,omitempty"`
		NameLocalised         string  `mapstructure:"Name_Localised,omitempty"`
		Quality               float64 `mapstructure:"Quality,omitempty"`
		StarSystem            string  `mapstructure:"StarSystem,omitempty"`
		StorageSlot           int     `mapstructure:"StorageSlot,omitempty"`
		TransferCost          int     `mapstructure:"TransferCost,omitempty"`
		TransferTime          int     `mapstructure:"TransferTime,omitempty"`
	} `mapstructure:"Items,omitempty"`
	MarketID    int       `mapstructure:"MarketID,omitempty"`
	StarSystem  string    `mapstructure:"StarSystem,omitempty"`
	StationName string    `mapstructure:"StationName,omitempty"`
	Event       string    `mapstructure:"event,omitempty"`
	Timestamp   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evStoredModules(ev *StoredModulesEvent) {
	return
}
