package file

import "time"

type LocationEvent struct {
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
	DistFromStarLs float64 `json:"DistFromStarLS,omitempty"`
	Docked         bool    `json:"Docked,omitempty"`
	Factions       []struct {
		ActiveStates []struct {
			State string `json:"State,omitempty"`
		} `json:"ActiveStates,omitempty"`
		Allegiance         string  `json:"Allegiance,omitempty"`
		FactionState       string  `json:"FactionState,omitempty"`
		Government         string  `json:"Government,omitempty"`
		HappiestSystem     bool    `json:"HappiestSystem,omitempty"`
		Happiness          string  `json:"Happiness,omitempty"`
		HappinessLocalised string  `json:"Happiness_Localised,omitempty"`
		HomeSystem         bool    `json:"HomeSystem,omitempty"`
		Influence          float64 `json:"Influence,omitempty"`
		MyReputation       float64 `json:"MyReputation,omitempty"`
		Name               string  `json:"Name,omitempty"`
		PendingStates      []struct {
			State string `json:"State,omitempty"`
			Trend int    `json:"Trend,omitempty"`
		} `json:"PendingStates,omitempty"`
		RecoveringStates []struct {
			State string `json:"State,omitempty"`
			Trend int    `json:"Trend,omitempty"`
		} `json:"RecoveringStates,omitempty"`
		SquadronFaction bool `json:"SquadronFaction,omitempty"`
	} `json:"Factions,omitempty"`
	InSrv             bool      `json:"InSRV,omitempty"`
	Latitude          float64   `json:"Latitude,omitempty"`
	Longitude         float64   `json:"Longitude,omitempty"`
	MarketID          int       `json:"MarketID,omitempty"`
	Multicrew         bool      `json:"Multicrew,omitempty"`
	OnFoot            bool      `json:"OnFoot,omitempty"`
	Population        int       `json:"Population,omitempty"`
	PowerplayState    string    `json:"PowerplayState,omitempty"`
	Powers            []string  `json:"Powers,omitempty"`
	StarPos           []float64 `json:"StarPos,omitempty"`
	StarSystem        string    `json:"StarSystem,omitempty"`
	StationAllegiance string    `json:"StationAllegiance,omitempty"`
	StationEconomies  []struct {
		Name          string  `json:"Name,omitempty"`
		NameLocalised string  `json:"Name_Localised,omitempty"`
		Proportion    float64 `json:"Proportion,omitempty"`
	} `json:"StationEconomies,omitempty"`
	StationEconomy          string `json:"StationEconomy,omitempty"`
	StationEconomyLocalised string `json:"StationEconomy_Localised,omitempty"`
	StationFaction          *struct {
		FactionState string `json:"FactionState,omitempty"`
		Name         string `json:"Name,omitempty"`
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
	SystemFaction              *struct {
		FactionState string `json:"FactionState,omitempty"`
		Name         string `json:"Name,omitempty"`
	} `json:"SystemFaction,omitempty"`
	SystemGovernment             string    `json:"SystemGovernment,omitempty"`
	SystemGovernmentLocalised    string    `json:"SystemGovernment_Localised,omitempty"`
	SystemSecondEconomy          string    `json:"SystemSecondEconomy,omitempty"`
	SystemSecondEconomyLocalised string    `json:"SystemSecondEconomy_Localised,omitempty"`
	SystemSecurity               string    `json:"SystemSecurity,omitempty"`
	SystemSecurityLocalised      string    `json:"SystemSecurity_Localised,omitempty"`
	Taxi                         bool      `json:"Taxi,omitempty"`
	Event                        string    `json:"event,omitempty"`
	Timestamp                    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLocation(ev *LocationEvent) {

	CurrentSystemName = ev.StarSystem

	return
}
