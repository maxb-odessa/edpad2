package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"fmt"
	"time"

	"github.com/danwakefield/fnmatch"
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

	var text string

	pd, ok := CurrentSystemPlanets[ev.BodyName]
	if !ok {
		return
	}

	for _, gen := range ev.Genuses {
		text += guessFlora(pd, gen.GenusLocalised, " |- ")
	}
	/*
		for _, sig := range ev.Signals {
			text += fmt.Sprintf(" | Sig: %s (%d)\n", sig.TypeLocalised, sig.Count)
		}
	*/
	if text != "" {
		text = `BIOs (valuable) at <i><b>` + ev.BodyName + `</b></i>` + "\n" + text

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

		slog.Debug(5, "SIGNALS: %s\n%+v", text, ev)
	}

	return
}

func guessFlora(pd *planetData, name string, prefix string) string {
	var text string

	for floraName, floraVar := range floras {

		// match something like "Aleoida" over "Aleoida Gravis"
		slog.Debug(15, "FLORA MATCH %s and %s\n", "*"+name+" *", floraName)
		if !fnmatch.Match("*"+name+" *", floraName, fnmatch.FNM_IGNORECASE) {
			continue
		}

		slog.Debug(15, "FLORA MATCHED! %s and %s\n", "*"+name+" *", floraName)
		if matches(pd.atmosphere, floraVar.atmos) &&
			matches(pd.class, floraVar.bodies) &&
			pd.gravityG >= floraVar.gravG[0] && pd.gravityG <= floraVar.gravG[1] &&
			pd.temperatureK >= floraVar.tempK[0] && pd.temperatureK <= floraVar.tempK[1] {
			text += fmt.Sprintf("%s%s, price %.2f Mil\n", prefix, floraName, floraVar.priceM)
		}

	}

	return text
}

type variant struct {
	priceM float32
	bodies []string
	atmos  []string
	gravG  []float64
	tempK  []float64
}

var floras = map[string]variant{
	"Aleoida Arcus": {
		priceM: 7.2,
		atmos:  []string{"*thin carbon dioxide*"},
		bodies: []string{"*high metal*", "*rocky body*"},
		tempK:  []float64{175.0, 180.0},
		gravG:  []float64{0.0, 0.27},
	},
	"Aleoida Coronamus": {
		priceM: 6.2,
		atmos:  []string{"*thin carbon dioxide*"},
		bodies: []string{"*high metal*", "*rocky body*"},
		tempK:  []float64{180.0, 190.0},
		gravG:  []float64{0.0, 0.27},
	},
	"Aleoida Gravis": {
		priceM: 12.9,
		atmos:  []string{"*thin carbon dioxide*"},
		bodies: []string{"*high metal*", "*rocky body*"},
		tempK:  []float64{190.0, 195.5},
		gravG:  []float64{0.0, 0.27},
	},
	"Bacterium Informem": {
		priceM: 8.4,
		atmos:  []string{"*nitrogen*"},
		bodies: []string{"*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 9999.0},
	},
	"Bacterium Nebulus": {
		priceM: 5.2,
		atmos:  []string{"*thin helium*"},
		bodies: []string{"*icy*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 9999.0},
	},
	"Bacterium Volu": {
		priceM: 7.7,
		atmos:  []string{"*oxygen*"},
		bodies: []string{"*high metal*", "*icy*", "*rocky ice*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 9999.0},
	},
	"Cactoida Vermis": {
		priceM: 16.2,
		atmos:  []string{"*water*"},
		bodies: []string{"*high metal*", "*rocky body*"},
		tempK:  []float64{165.0, 450.0},
		gravG:  []float64{0.0, 0.27},
	},
	"Clypeus Lacrimam": {
		priceM: 8.4,
		atmos:  []string{"*carbon dioxide*", "*water*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{190.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},
	"Clypeus Margaritus": {
		priceM: 11.8,
		atmos:  []string{"*carbon dioxide*", "*water*"},
		bodies: []string{"*high metal*"},
		tempK:  []float64{190.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Clypeus Speculumi (>2500Ls to star)": {
		priceM: 16.2,
		atmos:  []string{"*carbon dioxide*", "*water*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{190.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Concha Aureolas": {
		priceM: 7.7,
		atmos:  []string{"*ammonia*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Concha Biconcavis": {
		priceM: 19.0,
		atmos:  []string{"*nitrogen*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Electricae (All)": {
		priceM: 6.2,
		atmos:  []string{"*helium*", "*neon*", "*argon*"},
		bodies: []string{"*icy*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Fonticulua Fluctus": {
		priceM: 20.0,
		atmos:  []string{"*oxygen*"},
		bodies: []string{"*rocky ice*", "*icy*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Fonticulua Segmentatus": {
		priceM: 19.0,
		atmos:  []string{"*neon*"},
		bodies: []string{"*rocky ice*", "*icy*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Fonticulua Upupam": {
		priceM: 5.7,
		atmos:  []string{"*argon*rich*"},
		bodies: []string{"*icy*", "*rocky ice*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Frutexa Acus": {
		priceM: 7.7,
		atmos:  []string{"*carbon dioxide*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{0.0, 195.5},
		gravG:  []float64{0.0, 0.27},
	},

	"Frutexa Flammasis": {
		priceM: 10.3,
		atmos:  []string{"*ammonia*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Frutexa Sponsae": {
		priceM: 5.9,
		atmos:  []string{"*water*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Fumerola (All)": {
		priceM: 7.0,
		atmos:  []string{"*"},
		bodies: []string{"*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Osseus Discus (on rocks)": {
		priceM: 12.9,
		atmos:  []string{"*water*"},
		bodies: []string{"*high metal*", "*rocky body*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Osseus Pellebantus (on rocks)": {
		priceM: 9.7,
		atmos:  []string{"*carbon dioxide*"},
		bodies: []string{"*high metal*", "*rocky body*"},
		tempK:  []float64{180.0, 195.5},
		gravG:  []float64{0.0, 0.27},
	},

	"Recepta Conditivus": {
		priceM: 14.3,
		atmos:  []string{"*sulphur dioxide*"},
		bodies: []string{"*icy*", "*rocky ice*"},
		tempK:  []float64{132.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Recepta Deltahedronix": {
		priceM: 16.2,
		atmos:  []string{"*sulphur dioxide*"},
		bodies: []string{"*high metal*", "*rocky body*", "*icy*"},
		tempK:  []float64{132.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Recepta Umbrux": {
		priceM: 12.9,
		atmos:  []string{"*sulphur dioxide*"},
		bodies: []string{"*icy*", "*rocky*", "*ice*"},
		tempK:  []float64{132.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Stratum Cucumisis": {
		priceM: 16.2,
		atmos:  []string{"*thin carbon dioxide*", "*sulphur dioxide*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{190.0, 9999.0},
		gravG:  []float64{0.0, 9999.0},
	},

	"Stratum Tectonicas": {
		priceM: 19.0,
		atmos:  []string{"*thin*"},
		bodies: []string{"*high metal*"},
		tempK:  []float64{165.0, 9999.0},
		gravG:  []float64{0.0, 9999.0},
	},

	"Tubus (Cavas 11, Compagibus 7)": {
		priceM: 11.0,
		atmos:  []string{"*carbon dioxide*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{160.0, 190.5},
		gravG:  []float64{0.0, 0.15},
	},

	"Tubus Sororibus": {
		priceM: 5.7,
		atmos:  []string{"*ammonia*", "*carbon dioxide*"},
		bodies: []string{"*high metal*"},
		tempK:  []float64{160.0, 190.5},
		gravG:  []float64{0.0, 0.15},
	},

	"Tussock Capillum": {
		priceM: 7.0,
		atmos:  []string{"*argon*", "*methane*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Tussock Pennata": {
		priceM: 5.8,
		atmos:  []string{"*carbon dioxide*"},
		bodies: []string{"*high metal*", "*rocky body*"},
		tempK:  []float64{145.0, 155.5},
		gravG:  []float64{0.0, 0.27},
	},

	"Tussock Stigmasis": {
		priceM: 19.0,
		atmos:  []string{"*sulphur dioxide*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{0.0, 9999.0},
		gravG:  []float64{0.0, 0.27},
	},

	"Tussock Triticum": {
		priceM: 7.7,
		atmos:  []string{"*carbon dioxide*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{190.0, 195.5},
		gravG:  []float64{0.0, 0.27},
	},

	"Tussock Virgam": {
		priceM: 14.3,
		atmos:  []string{"*water*"},
		bodies: []string{"*rocky body*"},
		tempK:  []float64{0.0, 99999.9},
		gravG:  []float64{0.0, 0.27},
	},
}
