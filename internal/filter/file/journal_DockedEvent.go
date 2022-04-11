package file

import "time"

type DockedEvent struct {
	ActiveFine        bool    `mapstructure:"ActiveFine,omitempty"`
	DistFromStarLs    float64 `mapstructure:"DistFromStarLS,omitempty"`
	MarketID          int     `mapstructure:"MarketID,omitempty"`
	Multicrew         bool    `mapstructure:"Multicrew,omitempty"`
	StarSystem        string  `mapstructure:"StarSystem,omitempty"`
	StationAllegiance string  `mapstructure:"StationAllegiance,omitempty"`
	StationEconomies  []struct {
		Name          string  `mapstructure:"Name,omitempty"`
		NameLocalised string  `mapstructure:"Name_Localised,omitempty"`
		Proportion    float64 `mapstructure:"Proportion,omitempty"`
	} `mapstructure:"StationEconomies,omitempty"`
	StationEconomy          string `mapstructure:"StationEconomy,omitempty"`
	StationEconomyLocalised string `mapstructure:"StationEconomy_Localised,omitempty"`
	StationFaction          struct {
		FactionState string `mapstructure:"FactionState,omitempty"`
		Name         string `mapstructure:"Name,omitempty"`
	} `mapstructure:"StationFaction,omitempty"`
	StationGovernment          string    `mapstructure:"StationGovernment,omitempty"`
	StationGovernmentLocalised string    `mapstructure:"StationGovernment_Localised,omitempty"`
	StationName                string    `mapstructure:"StationName,omitempty"`
	StationServices            []string  `mapstructure:"StationServices,omitempty"`
	StationType                string    `mapstructure:"StationType,omitempty"`
	SystemAddress              int       `mapstructure:"SystemAddress,omitempty"`
	Taxi                       bool      `mapstructure:"Taxi,omitempty"`
	Wanted                     bool      `mapstructure:"Wanted,omitempty"`
	Event                      string    `mapstructure:"event,omitempty"`
	Timestamp                  time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evDocked(ev *DockedEvent) {
	return
}
