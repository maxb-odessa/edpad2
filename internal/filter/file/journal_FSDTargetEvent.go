package file

import (
	localDisplay "edpad2/internal/local/display"
	"edpad2/internal/router"
	"fmt"
	"time"
)

type FSDTargetEvent struct {
	Name                  string    `json:"Name,omitempty"`
	RemainingJumpsInRoute int       `json:"RemainingJumpsInRoute,omitempty"`
	StarClass             string    `json:"StarClass,omitempty"`
	SystemAddress         int       `json:"SystemAddress,omitempty"`
	Event                 string    `json:"event,omitempty"`
	Timestamp             time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFSDTarget(ev *FSDTargetEvent) {

	NextJumpSystem = ev.Name

	text := fmt.Sprintf(`<i>%s</i> (%s) >> <span color="white">%d</span> >> (%s) <i>%s</i>`,
		CurrentSystemName,
		CB(CurrentMainStarClass),
		ev.RemainingJumpsInRoute,
		CB(ev.StarClass),
		ev.Name)

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &localDisplay.Text{
			ViewPort:       localDisplay.VP_ROUTE,
			Text:           text,
			AppendText:     false,
			UpdateText:     true,
			Subtitle:       "",
			UpdateSubtitle: false,
		},
	}

	return
}
