package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"time"

	"github.com/maxb-odessa/slog"
)

type FSSAllBodiesFoundEvent struct {
	Count         int       `mapstructure:"Count,omitempty"`
	SystemAddress int       `mapstructure:"SystemAddress,omitempty"`
	SystemName    string    `mapstructure:"SystemName,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evFSSAllBodiesFound(ev *FSSAllBodiesFoundEvent) {

	for pname, pdata := range CurrentSystemPlanets {
		var text string

		for _, bio := range pdata.bios {
			text += guessFlora(pdata, bio, " ?- ")
		}

		if text == "" {
			continue
		}

		text = `BIOs (valuable) at <i><b>` + pname + `</b></i>` + "\n" + text

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

		slog.Debug(5, "BIOS: %s\n%+v", text, ev)
	}

	return
}
