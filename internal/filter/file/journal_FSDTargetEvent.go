package file

import (
	"edpad2/internal/local/display"
	localDisplay "edpad2/internal/local/display"
	"edpad2/internal/router"
	"fmt"
	"time"
)

type FSDTargetEvent struct {
	Name                  string    `mapstructure:"Name,omitempty"`
	RemainingJumpsInRoute int       `mapstructure:"RemainingJumpsInRoute,omitempty"`
	StarClass             string    `mapstructure:"StarClass,omitempty"`
	SystemAddress         int       `mapstructure:"SystemAddress,omitempty"`
	Event                 string    `mapstructure:"event,omitempty"`
	Timestamp             time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evFSDTarget(ev *FSDTargetEvent) {

	NextJumpSystem = ev.Name
	NextJumpStarClass = ev.StarClass

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

	return
}
