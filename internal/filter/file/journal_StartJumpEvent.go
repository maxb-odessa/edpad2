package file

import (
	localDisplay "edpad2/internal/local/display"
	"edpad2/internal/router"
	"time"
)

type StartJumpEvent struct {
	JumpType      string    `json:"JumpType,omitempty"`
	StarClass     string    `json:"StarClass,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evStartJump(ev *StartJumpEvent) {

	if ev.JumpType != "Hyperspace" {
		return
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
