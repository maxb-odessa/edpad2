package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"fmt"
	"time"

	"github.com/maxb-odessa/slog"
)

type SAASignalsFoundEvent struct {
	BodyID   int    `mapstructure:"BodyID,omitempty"`
	BodyName string `mapstructure:"BodyName,omitempty"`
	Signals  []struct {
		Count         int    `mapstructure:"Count,omitempty"`
		Type          string `mapstructure:"Type,omitempty"`
		TypeLocalised string `mapstructure:"Type_Localised,omitempty"`
	} `mapstructure:"Signals,omitempty"`
	Genuses []struct {
		Genus          string `mapstructure:"Genus,omitempty"`
		GenusLocalised string `mapstructure:"Genus_Localised,omitempty"`
	} `mapstructure:"Genuses,omitempty"`

	SystemAddress int       `mapstructure:"SystemAddress,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

// DSS scan complete
func (h *handler) evSAASignalsFound(ev *SAASignalsFoundEvent) {

	pd, ok := CurrentSystemPlanets[ev.BodyName]
	if !ok {
		return
	}

	var text string

	for _, gen := range ev.Genuses {
		f := guessFlora(pd, gen.GenusLocalised)
		if f != "" {
			text += fmt.Sprintf(" | Bio: %s\n", f)
		}
	}

	for _, sig := range ev.Signals {
		text += fmt.Sprintf(" | Sig: %s (%d)\n", sig.TypeLocalised, sig.Count)
	}

	if text != "" {
		text = `<i>` + ev.BodyName + `</i>` + "\n" + text

		h.connector.ToRouterCh <- &router.Message{
			Dst: router.LocalDisplay,
			Data: &display.Text{
				ViewPort:       display.VP_SIGNALS,
				Text:           text,
				AppendText:     true,
				UpdateText:     true,
				Subtitle:       `[!]`,
				UpdateSubtitle: true,
			},
		}

		slog.Debug(5, "SIGNAL: %s\n%+v", text, ev)
	}

	return
}

func guessFlora(pd *planetData, name string) string {

	for floraName, floraVariants := range floras {

		if floraName != name {
			continue
		}

		for varName, varData := range floraVariants {

			if matches(pd.atmosphere, varData.atmos) &&
				matches(pd.class, varData.bodies) &&
				pd.gravityG >= varData.gravG[0] && pd.gravityG <= varData.gravG[1] &&
				pd.temperatureK >= varData.tempK[0] && pd.temperatureK <= varData.tempK[1] {
				return fmt.Sprintf("%s %s, price %d", floraName, varName, varData.price)

			}
		}

	}

	return ""
}

type variant struct {
	price  int
	bodies []string
	atmos  []string
	gravG  []float64
	tempK  []float64
}

var floras = map[string]map[string]variant{
	"Aleoida": {
		"arcus": {
			price:  379300,
			bodies: []string{"rocky body*", "*high metal*"},
			atmos:  []string{"*thin carbon dioxide*"},
			gravG:  []float64{0.0, 0.27},
			tempK:  []float64{175.0, 180.0},
		},
		"TODO": {
			price:  0,
			bodies: []string{"*"},
			atmos:  []string{"*"},
			gravG:  []float64{0.0, 99999.0},
			tempK:  []float64{0.0, 9999.0},
		},
	},
}
