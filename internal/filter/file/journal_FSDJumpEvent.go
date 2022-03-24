package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"fmt"
	"time"
)

type FSDJumpEvent struct {
	Body      string `json:"Body,omitempty"`
	BodyID    int    `json:"BodyID,omitempty"`
	BodyType  string `json:"BodyType,omitempty"`
	BoostUsed int    `json:"BoostUsed,omitempty"`
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
	Factions []struct {
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
	FuelLevel              float64   `json:"FuelLevel,omitempty"`
	FuelUsed               float64   `json:"FuelUsed,omitempty"`
	JumpDist               float64   `json:"JumpDist,omitempty"`
	Multicrew              bool      `json:"Multicrew,omitempty"`
	Population             int       `json:"Population,omitempty"`
	PowerplayState         string    `json:"PowerplayState,omitempty"`
	Powers                 []string  `json:"Powers,omitempty"`
	StarPos                []float64 `json:"StarPos,omitempty"`
	StarSystem             string    `json:"StarSystem,omitempty"`
	SystemAddress          int       `json:"SystemAddress,omitempty"`
	SystemAllegiance       string    `json:"SystemAllegiance,omitempty"`
	SystemEconomy          string    `json:"SystemEconomy,omitempty"`
	SystemEconomyLocalised string    `json:"SystemEconomy_Localised,omitempty"`
	SystemFaction          *struct {
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

func (h *handler) evFSDJump(ev *FSDJumpEvent) {

	CurrentSystemName = ev.StarSystem // really neede? already set in StartJump event

	resetCurrentSystemStars()
	resetCurrentSystemPlanets()
	resetCurrentSystemSignals()

	text := fmt.Sprintf(`<span size="large">`+
		"\n\nSystem: %s\n\n"+
		`<i>`+
		"Jump distance: %3.2f ly\n"+
		"   Fuel level: %3.2f tons\n"+
		"    Fuel used: %3.2f tons\n"+
		`</i>`+
		`</span>`,
		ev.StarSystem,
		ev.JumpDist,
		ev.FuelLevel,
		ev.FuelUsed)

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_SYSTEM,
			Text:           text,
			AppendText:     false,
			UpdateText:     true,
			Subtitle:       "",
			UpdateSubtitle: true,
		},
	}

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_PLANETS,
			Text:           "",
			AppendText:     false,
			UpdateText:     true,
			Subtitle:       "",
			UpdateSubtitle: true,
		},
	}

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_SIGNALS,
			Text:           "",
			AppendText:     false,
			UpdateText:     true,
			Subtitle:       "",
			UpdateSubtitle: true,
		},
	}

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_SRV,
			Text:           "",
			AppendText:     false,
			UpdateText:     true,
			Subtitle:       "",
			UpdateSubtitle: true,
		},
	}

	CurrentMainStarName = ev.Body
	CurrentSystemName = ev.StarSystem

	return
}
