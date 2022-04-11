package file

import "time"

type LoadGameEvent struct {
	Commander     string    `mapstructure:"Commander,omitempty"`
	Credits       int       `mapstructure:"Credits,omitempty"`
	Fid           string    `mapstructure:"FID,omitempty"`
	FuelCapacity  float64   `mapstructure:"FuelCapacity,omitempty"`
	FuelLevel     float64   `mapstructure:"FuelLevel,omitempty"`
	GameMode      string    `mapstructure:"GameMode,omitempty"`
	Horizons      bool      `mapstructure:"Horizons,omitempty"`
	Loan          int       `mapstructure:"Loan,omitempty"`
	Odyssey       bool      `mapstructure:"Odyssey,omitempty"`
	Ship          string    `mapstructure:"Ship,omitempty"`
	ShipID        int       `mapstructure:"ShipID,omitempty"`
	ShipIdent     string    `mapstructure:"ShipIdent,omitempty"`
	ShipName      string    `mapstructure:"ShipName,omitempty"`
	ShipLocalised string    `mapstructure:"Ship_Localised,omitempty"`
	StartLanded   bool      `mapstructure:"StartLanded,omitempty"`
	Build         string    `mapstructure:"build,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Gameversion   string    `mapstructure:"gameversion,omitempty"`
	Language      string    `mapstructure:"language,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evLoadGame(ev *LoadGameEvent) {
	return
}
