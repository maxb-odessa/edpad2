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

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &localDisplay.Text{
			ViewPort: localDisplay.VP_SYSTEM,
			Text:     "TODO",
			// TODO Text:           `<i><span size="x-large">` + "\n\nJumping to: \n\n" + ev.StarSystem + " (" + CB(ev.StarClass) + ")" + `</span></i>`,
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
