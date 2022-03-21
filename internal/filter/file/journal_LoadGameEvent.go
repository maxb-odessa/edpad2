package file

import "time"

type LoadGameEvent struct {
	Commander     string    `json:"Commander,omitempty"`
	Credits       int       `json:"Credits,omitempty"`
	Fid           string    `json:"FID,omitempty"`
	FuelCapacity  float64   `json:"FuelCapacity,omitempty"`
	FuelLevel     float64   `json:"FuelLevel,omitempty"`
	GameMode      string    `json:"GameMode,omitempty"`
	Horizons      bool      `json:"Horizons,omitempty"`
	Loan          int       `json:"Loan,omitempty"`
	Odyssey       bool      `json:"Odyssey,omitempty"`
	Ship          string    `json:"Ship,omitempty"`
	ShipID        int       `json:"ShipID,omitempty"`
	ShipIdent     string    `json:"ShipIdent,omitempty"`
	ShipName      string    `json:"ShipName,omitempty"`
	ShipLocalised string    `json:"Ship_Localised,omitempty"`
	StartLanded   bool      `json:"StartLanded,omitempty"`
	Build         string    `json:"build,omitempty"`
	Event         string    `json:"event,omitempty"`
	Gameversion   string    `json:"gameversion,omitempty"`
	Language      string    `json:"language,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLoadGame(ev *LoadGameEvent) {
	return
}
