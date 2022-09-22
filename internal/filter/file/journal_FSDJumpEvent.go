package file

import (
	"edpad2/internal/local/display"
	localDisplay "edpad2/internal/local/display"
	"edpad2/internal/local/sound"
	"edpad2/internal/router"
	"fmt"
	"time"
)

type FSDJumpEvent struct {
	Body      string `mapstructure:"Body,omitempty"`
	BodyID    int    `mapstructure:"BodyID,omitempty"`
	BodyType  string `mapstructure:"BodyType,omitempty"`
	BoostUsed int    `mapstructure:"BoostUsed,omitempty"`
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
	Factions []struct {
		ActiveStates []struct {
			State string `mapstructure:"State,omitempty"`
		} `mapstructure:"ActiveStates,omitempty"`
		Allegiance         string  `mapstructure:"Allegiance,omitempty"`
		FactionState       string  `mapstructure:"FactionState,omitempty"`
		Government         string  `mapstructure:"Government,omitempty"`
		HappiestSystem     bool    `mapstructure:"HappiestSystem,omitempty"`
		Happiness          string  `mapstructure:"Happiness,omitempty"`
		HappinessLocalised string  `mapstructure:"Happiness_Localised,omitempty"`
		HomeSystem         bool    `mapstructure:"HomeSystem,omitempty"`
		Influence          float64 `mapstructure:"Influence,omitempty"`
		MyReputation       float64 `mapstructure:"MyReputation,omitempty"`
		Name               string  `mapstructure:"Name,omitempty"`
		PendingStates      []struct {
			State string `mapstructure:"State,omitempty"`
			Trend int    `mapstructure:"Trend,omitempty"`
		} `mapstructure:"PendingStates,omitempty"`
		RecoveringStates []struct {
			State string `mapstructure:"State,omitempty"`
			Trend int    `mapstructure:"Trend,omitempty"`
		} `mapstructure:"RecoveringStates,omitempty"`
		SquadronFaction bool `mapstructure:"SquadronFaction,omitempty"`
	} `mapstructure:"Factions,omitempty"`
	FuelLevel              float64   `mapstructure:"FuelLevel,omitempty"`
	FuelUsed               float64   `mapstructure:"FuelUsed,omitempty"`
	JumpDist               float64   `mapstructure:"JumpDist,omitempty"`
	Multicrew              bool      `mapstructure:"Multicrew,omitempty"`
	Population             int       `mapstructure:"Population,omitempty"`
	PowerplayState         string    `mapstructure:"PowerplayState,omitempty"`
	Powers                 []string  `mapstructure:"Powers,omitempty"`
	StarPos                []float64 `mapstructure:"StarPos,omitempty"`
	StarSystem             string    `mapstructure:"StarSystem,omitempty"`
	SystemAddress          int       `mapstructure:"SystemAddress,omitempty"`
	SystemAllegiance       string    `mapstructure:"SystemAllegiance,omitempty"`
	SystemEconomy          string    `mapstructure:"SystemEconomy,omitempty"`
	SystemEconomyLocalised string    `mapstructure:"SystemEconomy_Localised,omitempty"`
	SystemFaction          *struct {
		FactionState string `mapstructure:"FactionState,omitempty"`
		Name         string `mapstructure:"Name,omitempty"`
	} `mapstructure:"SystemFaction,omitempty"`
	SystemGovernment             string    `mapstructure:"SystemGovernment,omitempty"`
	SystemGovernmentLocalised    string    `mapstructure:"SystemGovernment_Localised,omitempty"`
	SystemSecondEconomy          string    `mapstructure:"SystemSecondEconomy,omitempty"`
	SystemSecondEconomyLocalised string    `mapstructure:"SystemSecondEconomy_Localised,omitempty"`
	SystemSecurity               string    `mapstructure:"SystemSecurity,omitempty"`
	SystemSecurityLocalised      string    `mapstructure:"SystemSecurity_Localised,omitempty"`
	Taxi                         bool      `mapstructure:"Taxi,omitempty"`
	Event                        string    `mapstructure:"event,omitempty"`
	Timestamp                    time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evFSDJump(ev *FSDJumpEvent) {

	CurrentSystemName = ev.StarSystem // really needed? already set in StartJump event

	if len(CurrentMainStarClass) > 1 {
		switch CurrentMainStarClass[0:1] {
		case "N", "D", "H":
			h.connector.ToRouterCh <- &router.Message{
				Dst: router.LocalSound,
				Data: &sound.Track{
					Id: sound.TONE,
				},
			}
		}
	}

	// no more jumps in route
	if CurrentSystemName == NextJumpSystem {
		h.connector.ToRouterCh <- &router.Message{
			Dst: router.LocalDisplay,
			Data: &localDisplay.Text{
				ViewPort:       localDisplay.VP_ROUTE,
				Text:           `<i>` + CurrentSystemName + `</i>`,
				AppendText:     false,
				UpdateText:     true,
				Subtitle:       "",
				UpdateSubtitle: false,
			},
		}
	}

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

	FuelLevel = ev.FuelLevel

	alertText := ""
	if ev.FuelLevel < 12.0 {

		alertText = `<span color="yellow">`

		if ev.FuelLevel < 6.0 {
			alertText = `<span color="red">`
		}

		alertText += fmt.Sprintf("Alert! Low fuel level (%.2f tons)", ev.FuelLevel)
		alertText += `</span>`

		h.connector.ToRouterCh <- &router.Message{
			Dst: router.LocalSound,
			Data: &sound.Track{
				Id:     sound.DROP,
				Repeat: 3,
			},
		}

	}

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_SYSTEM,
			Text:           text,
			AppendText:     false,
			UpdateText:     true,
			Subtitle:       "",
			UpdateSubtitle: true,
			SetVisible:     true,
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
			Text:           "\n",
			AppendText:     false,
			UpdateText:     true,
			Subtitle:       "",
			UpdateSubtitle: true,
		},
	}

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_LOGS,
			Text:           "\n<u>" + CurrentSystemName + "</u>\n",
			AppendText:     true,
			UpdateText:     true,
			Subtitle:       "",
			UpdateSubtitle: true,
		},
	}

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:   display.VP_INFO,
			Text:       alertText,
			AppendText: false,
			UpdateText: true,
		},
	}

	CurrentMainStarName = ev.Body
	CurrentSystemName = ev.StarSystem

	return
}
