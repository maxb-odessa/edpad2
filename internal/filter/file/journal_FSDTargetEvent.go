package file

import (
	"edpad2/internal/local/display"
	localDisplay "edpad2/internal/local/display"
	"edpad2/internal/local/sound"
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

	csclass, cscolor := CB(CurrentMainStarClass)
	jtclass, jtcolor := CB(ev.StarClass)

	currsys := `<i>` + CurrentSystemName + ` </i><span fgcolor="` + cscolor + `">` + `(` + csclass + `)</span>`
	tgtsys := `<span fgcolor="` + jtcolor + `">` + `(` + jtclass + `)</span>` + ` <i>` + NextJumpSystem + `</i>`

	text := fmt.Sprintf(`%s%s<span color="white">%d</span>%s%s`,
		currsys,
		`<span color="gray">`+display.HEAVYDASHEDRIGHTARROW+`</span>`,
		ev.RemainingJumpsInRoute,
		`<span color="gray">`+display.HEAVYDASHEDRIGHTARROW+`</span>`,
		tgtsys)

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &localDisplay.Text{
			ViewPort:   localDisplay.VP_ROUTE,
			Text:       text,
			AppendText: false,
			UpdateText: true,
		},
	}

	switch jtclass[0:1] {
	case "N", "D", "H":
		h.connector.ToRouterCh <- &router.Message{
			Dst: router.LocalSound,
			Data: &sound.Track{
				Id: sound.ALARM,
			},
		}
	}

	return
}
