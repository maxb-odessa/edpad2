package file

import "time"

type CarrierJumpEvent struct {
	Body      string `mapstructure:"Body,omitempty"`
	BodyID    int    `mapstructure:"BodyID,omitempty"`
	BodyType  string `mapstructure:"BodyType,omitempty"`
	Conflicts []struct {
		Faction1 struct {
			Name    string `mapstructure:"Name,omitempty"`
			Stake   string `mapstructure:"Stake,omitempty"`
			WonDays int    `mapstructure:"WonDays,omitempty"`
		} `mapstructure:"Faction1,omitempty"`
		Faction2 struct {
			Name    string `mapstructure:"Name,omitempty"`
			Stake   string `mapstructure:"Stake,omitempty"`
			WonDays int    `mapstructure:"WonDays,omitempty"`
		} `mapstructure:"Faction2,omitempty"`
		Status  string `mapstructure:"Status,omitempty"`
		WarType string `mapstructure:"WarType,omitempty"`
	} `mapstructure:"Conflicts,omitempty"`
	Docked   bool `mapstructure:"Docked,omitempty"`
	Factions []struct {
		ActiveStates []struct {
			State string `mapstructure:"State,omitempty"`
		} `mapstructure:"ActiveStates,omitempty"`
		Allegiance         string  `mapstructure:"Allegiance,omitempty"`
		FactionState       string  `mapstructure:"FactionState,omitempty"`
		Government         string  `mapstructure:"Government,omitempty"`
		Happiness          string  `mapstructure:"Happiness,omitempty"`
		HappinessLocalised string  `mapstructure:"Happiness_Localised,omitempty"`
		Influence          float64 `mapstructure:"Influence,omitempty"`
		MyReputation       float64 `mapstructure:"MyReputation,omitempty"`
		Name               string  `mapstructure:"Name,omitempty"`
	} `mapstructure:"Factions,omitempty"`
	MarketID         int       `mapstructure:"MarketID,omitempty"`
	Population       int       `mapstructure:"Population,omitempty"`
	StarPos          []float64 `mapstructure:"StarPos,omitempty"`
	StarSystem       string    `mapstructure:"StarSystem,omitempty"`
	StationEconomies []struct {
		Name          string  `mapstructure:"Name,omitempty"`
		NameLocalised string  `mapstructure:"Name_Localised,omitempty"`
		Proportion    float64 `mapstructure:"Proportion,omitempty"`
	} `mapstructure:"StationEconomies,omitempty"`
	StationEconomy          string `mapstructure:"StationEconomy,omitempty"`
	StationEconomyLocalised string `mapstructure:"StationEconomy_Localised,omitempty"`
	StationFaction          struct {
		Name string `mapstructure:"Name,omitempty"`
	} `mapstructure:"StationFaction,omitempty"`
	StationGovernment          string   `mapstructure:"StationGovernment,omitempty"`
	StationGovernmentLocalised string   `mapstructure:"StationGovernment_Localised,omitempty"`
	StationName                string   `mapstructure:"StationName,omitempty"`
	StationServices            []string `mapstructure:"StationServices,omitempty"`
	StationType                string   `mapstructure:"StationType,omitempty"`
	SystemAddress              int      `mapstructure:"SystemAddress,omitempty"`
	SystemAllegiance           string   `mapstructure:"SystemAllegiance,omitempty"`
	SystemEconomy              string   `mapstructure:"SystemEconomy,omitempty"`
	SystemEconomyLocalised     string   `mapstructure:"SystemEconomy_Localised,omitempty"`
	SystemFaction              struct {
		FactionState string `mapstructure:"FactionState,omitempty"`
		Name         string `mapstructure:"Name,omitempty"`
	} `mapstructure:"SystemFaction,omitempty"`
	SystemGovernment             string    `mapstructure:"SystemGovernment,omitempty"`
	SystemGovernmentLocalised    string    `mapstructure:"SystemGovernment_Localised,omitempty"`
	SystemSecondEconomy          string    `mapstructure:"SystemSecondEconomy,omitempty"`
	SystemSecondEconomyLocalised string    `mapstructure:"SystemSecondEconomy_Localised,omitempty"`
	SystemSecurity               string    `mapstructure:"SystemSecurity,omitempty"`
	SystemSecurityLocalised      string    `mapstructure:"SystemSecurity_Localised,omitempty"`
	Event                        string    `mapstructure:"event,omitempty"`
	Timestamp                    time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCarrierJump(ev *CarrierJumpEvent) {
	return
}
