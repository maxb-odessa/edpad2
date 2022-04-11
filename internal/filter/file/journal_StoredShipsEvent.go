package file

import "time"

type StoredShipsEvent struct {
	MarketID  int `mapstructure:"MarketID,omitempty"`
	ShipsHere []struct {
		Hot               bool   `mapstructure:"Hot,omitempty"`
		Name              string `mapstructure:"Name,omitempty"`
		ShipID            int    `mapstructure:"ShipID,omitempty"`
		ShipType          string `mapstructure:"ShipType,omitempty"`
		ShipTypeLocalised string `mapstructure:"ShipType_Localised,omitempty"`
		Value             int    `mapstructure:"Value,omitempty"`
	} `mapstructure:"ShipsHere,omitempty"`
	ShipsRemote []struct {
		Hot               bool   `mapstructure:"Hot,omitempty"`
		InTransit         bool   `mapstructure:"InTransit,omitempty"`
		Name              string `mapstructure:"Name,omitempty"`
		ShipID            int    `mapstructure:"ShipID,omitempty"`
		ShipMarketID      int    `mapstructure:"ShipMarketID,omitempty"`
		ShipType          string `mapstructure:"ShipType,omitempty"`
		ShipTypeLocalised string `mapstructure:"ShipType_Localised,omitempty"`
		StarSystem        string `mapstructure:"StarSystem,omitempty"`
		TransferPrice     int    `mapstructure:"TransferPrice,omitempty"`
		TransferTime      int    `mapstructure:"TransferTime,omitempty"`
		Value             int    `mapstructure:"Value,omitempty"`
	} `mapstructure:"ShipsRemote,omitempty"`
	StarSystem  string    `mapstructure:"StarSystem,omitempty"`
	StationName string    `mapstructure:"StationName,omitempty"`
	Event       string    `mapstructure:"event,omitempty"`
	Timestamp   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evStoredShips(ev *StoredShipsEvent) {
	return
}
