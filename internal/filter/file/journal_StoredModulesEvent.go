package file

import "time"

type StoredModulesEvent struct {
	Items []struct {
		BuyPrice              int     `json:"BuyPrice,omitempty"`
		EngineerModifications string  `json:"EngineerModifications,omitempty"`
		Hot                   bool    `json:"Hot,omitempty"`
		InTransit             bool    `json:"InTransit,omitempty"`
		Level                 int     `json:"Level,omitempty"`
		MarketID              int     `json:"MarketID,omitempty"`
		Name                  string  `json:"Name,omitempty"`
		NameLocalised         string  `json:"Name_Localised,omitempty"`
		Quality               float64 `json:"Quality,omitempty"`
		StarSystem            string  `json:"StarSystem,omitempty"`
		StorageSlot           int     `json:"StorageSlot,omitempty"`
		TransferCost          int     `json:"TransferCost,omitempty"`
		TransferTime          int     `json:"TransferTime,omitempty"`
	} `json:"Items,omitempty"`
	MarketID    int       `json:"MarketID,omitempty"`
	StarSystem  string    `json:"StarSystem,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evStoredModules(ev *StoredModulesEvent) {
	return
}
