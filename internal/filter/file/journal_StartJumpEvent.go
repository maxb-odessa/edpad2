package file

import (
	localDisplay "edpad2/internal/local/display"
	"edpad2/internal/local/sound"
	"edpad2/internal/router"
	"time"
)

type StartJumpEvent struct {
	JumpType      string    `mapstructure:"JumpType,omitempty"`
	StarClass     string    `mapstructure:"StarClass,omitempty"`
	StarSystem    string    `mapstructure:"StarSystem,omitempty"`
	SystemAddress int       `mapstructure:"SystemAddress,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evStartJump(ev *StartJumpEvent) {

	if ev.JumpType != "Hyperspace" {
		return
	}

	switch ev.StarClass[0:1] {
	case "N", "D", "H":
		h.connector.ToRouterCh <- &router.Message{
			Dst: router.LocalSound,
			Data: &sound.Track{
				Id: sound.WARN,
			},
		}
	}

	sname, scolor := CB(ev.StarClass)

	text := `<i><span size="x-large">` + "\n\nJumping to: \n\n" + ev.StarSystem +
		` <span fgcolor="` + scolor + `">(` + sname + `)</span>` +
		`</span></i>`

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &localDisplay.Text{
			ViewPort:       localDisplay.VP_SYSTEM,
			Text:           text,
			AppendText:     false,
			UpdateText:     true,
			Subtitle:       "",
			UpdateSubtitle: true,
		},
	}

	CurrentSystemName = ev.StarSystem
	CurrentMainStarClass = ev.StarClass

	return
}
