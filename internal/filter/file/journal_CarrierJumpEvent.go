package file

import "time"

type CarrierJumpEvent struct {
	Body      string `json:"Body,omitempty"`
	BodyID    int    `json:"BodyID,omitempty"`
	BodyType  string `json:"BodyType,omitempty"`
	Conflicts []struct {
		Faction1 struct {
			Name    string `json:"Name,omitempty"`
			Stake   string `json:"Stake,omitempty"`
			WonDays int    `json:"WonDays,omitempty"`
		} `json:"Faction1,omitempty"`
		Faction2 struct {
			Name    string `json:"Name,omitempty"`
			Stake   string `json:"Stake,omitempty"`
			WonDays int    `json:"WonDays,omitempty"`
		} `json:"Faction2,omitempty"`
		Status  string `json:"Status,omitempty"`
		WarType string `json:"WarType,omitempty"`
	} `json:"Conflicts,omitempty"`
	Docked   bool `json:"Docked,omitempty"`
	Factions []struct {
		ActiveStates []struct {
			State string `json:"State,omitempty"`
		} `json:"ActiveStates,omitempty"`
		Allegiance         string  `json:"Allegiance,omitempty"`
		FactionState       string  `json:"FactionState,omitempty"`
		Government         string  `json:"Government,omitempty"`
		Happiness          string  `json:"Happiness,omitempty"`
		HappinessLocalised string  `json:"Happiness_Localised,omitempty"`
		Influence          float64 `json:"Influence,omitempty"`
		MyReputation       float64 `json:"MyReputation,omitempty"`
		Name               string  `json:"Name,omitempty"`
	} `json:"Factions,omitempty"`
	MarketID         int       `json:"MarketID,omitempty"`
	Population       int       `json:"Population,omitempty"`
	StarPos          []float64 `json:"StarPos,omitempty"`
	StarSystem       string    `json:"StarSystem,omitempty"`
	StationEconomies []struct {
		Name          string  `json:"Name,omitempty"`
		NameLocalised string  `json:"Name_Localised,omitempty"`
		Proportion    float64 `json:"Proportion,omitempty"`
	} `json:"StationEconomies,omitempty"`
	StationEconomy          string `json:"StationEconomy,omitempty"`
	StationEconomyLocalised string `json:"StationEconomy_Localised,omitempty"`
	StationFaction          struct {
		Name string `json:"Name,omitempty"`
	} `json:"StationFaction,omitempty"`
	StationGovernment          string   `json:"StationGovernment,omitempty"`
	StationGovernmentLocalised string   `json:"StationGovernment_Localised,omitempty"`
	StationName                string   `json:"StationName,omitempty"`
	StationServices            []string `json:"StationServices,omitempty"`
	StationType                string   `json:"StationType,omitempty"`
	SystemAddress              int      `json:"SystemAddress,omitempty"`
	SystemAllegiance           string   `json:"SystemAllegiance,omitempty"`
	SystemEconomy              string   `json:"SystemEconomy,omitempty"`
	SystemEconomyLocalised     string   `json:"SystemEconomy_Localised,omitempty"`
	SystemFaction              struct {
		FactionState string `json:"FactionState,omitempty"`
		Name         string `json:"Name,omitempty"`
	} `json:"SystemFaction,omitempty"`
	SystemGovernment             string    `json:"SystemGovernment,omitempty"`
	SystemGovernmentLocalised    string    `json:"SystemGovernment_Localised,omitempty"`
	SystemSecondEconomy          string    `json:"SystemSecondEconomy,omitempty"`
	SystemSecondEconomyLocalised string    `json:"SystemSecondEconomy_Localised,omitempty"`
	SystemSecurity               string    `json:"SystemSecurity,omitempty"`
	SystemSecurityLocalised      string    `json:"SystemSecurity_Localised,omitempty"`
	Event                        string    `json:"event,omitempty"`
	Timestamp                    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCarrierJump(ev *CarrierJumpEvent) {
	return
}
