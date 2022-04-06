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

	csname, cscolor := CB(CurrentMainStarClass)
	jtname, jtcolor := CB(ev.StarClass)

	currsys := `<i>` + CurrentSystemName + `</i> <span fgcolor="` + cscolor + `">` + `(` + csname + `)</span>`
	tgtsys := `<span fgcolor="` + jtcolor + `">` + `(` + jtname + `)</span>` + `<i>` + NextJumpSystem + `</i>`

	text := fmt.Sprintf(`%s &gt;&gt; <b><span color="white">%d</span></b> &gt;&gt; %s`,
		currsys,
		ev.RemainingJumpsInRoute,
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
