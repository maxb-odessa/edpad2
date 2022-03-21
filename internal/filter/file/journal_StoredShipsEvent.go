package file

import "time"

type StoredShipsEvent struct {
	MarketID  int `json:"MarketID,omitempty"`
	ShipsHere []struct {
		Hot               bool   `json:"Hot,omitempty"`
		Name              string `json:"Name,omitempty"`
		ShipID            int    `json:"ShipID,omitempty"`
		ShipType          string `json:"ShipType,omitempty"`
		ShipTypeLocalised string `json:"ShipType_Localised,omitempty"`
		Value             int    `json:"Value,omitempty"`
	} `json:"ShipsHere,omitempty"`
	ShipsRemote []struct {
		Hot               bool   `json:"Hot,omitempty"`
		InTransit         bool   `json:"InTransit,omitempty"`
		Name              string `json:"Name,omitempty"`
		ShipID            int    `json:"ShipID,omitempty"`
		ShipMarketID      int    `json:"ShipMarketID,omitempty"`
		ShipType          string `json:"ShipType,omitempty"`
		ShipTypeLocalised string `json:"ShipType_Localised,omitempty"`
		StarSystem        string `json:"StarSystem,omitempty"`
		TransferPrice     int    `json:"TransferPrice,omitempty"`
		TransferTime      int    `json:"TransferTime,omitempty"`
		Value             int    `json:"Value,omitempty"`
	} `json:"ShipsRemote,omitempty"`
	StarSystem  string    `json:"StarSystem,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evStoredShips(ev *StoredShipsEvent) {
	return
}
