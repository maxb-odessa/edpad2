package file

import "time"

type DockedEvent struct {
	ActiveFine        bool    `json:"ActiveFine,omitempty"`
	DistFromStarLs    float64 `json:"DistFromStarLS,omitempty"`
	MarketID          int     `json:"MarketID,omitempty"`
	Multicrew         bool    `json:"Multicrew,omitempty"`
	StarSystem        string  `json:"StarSystem,omitempty"`
	StationAllegiance string  `json:"StationAllegiance,omitempty"`
	StationEconomies  []struct {
		Name          string  `json:"Name,omitempty"`
		NameLocalised string  `json:"Name_Localised,omitempty"`
		Proportion    float64 `json:"Proportion,omitempty"`
	} `json:"StationEconomies,omitempty"`
	StationEconomy          string `json:"StationEconomy,omitempty"`
	StationEconomyLocalised string `json:"StationEconomy_Localised,omitempty"`
	StationFaction          struct {
		FactionState string `json:"FactionState,omitempty"`
		Name         string `json:"Name,omitempty"`
	} `json:"StationFaction,omitempty"`
	StationGovernment          string    `json:"StationGovernment,omitempty"`
	StationGovernmentLocalised string    `json:"StationGovernment_Localised,omitempty"`
	StationName                string    `json:"StationName,omitempty"`
	StationServices            []string  `json:"StationServices,omitempty"`
	StationType                string    `json:"StationType,omitempty"`
	SystemAddress              int       `json:"SystemAddress,omitempty"`
	Taxi                       bool      `json:"Taxi,omitempty"`
	Wanted                     bool      `json:"Wanted,omitempty"`
	Event                      string    `json:"event,omitempty"`
	Timestamp                  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDocked(ev *DockedEvent) {
	return
}
