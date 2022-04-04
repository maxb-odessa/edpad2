package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"fmt"
	"time"
)

type MaterialCollectedEvent struct {
	Category      string    `json:"Category,omitempty"`
	Count         int       `json:"Count,omitempty"`
	Name          string    `json:"Name,omitempty"`
	NameLocalised string    `json:"Name_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMaterialCollected(ev *MaterialCollectedEvent) {

	text := fmt.Sprintf("Material collected: %s\n --> %s (%d)\n", ev.Category, ev.Name, ev.Count)

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_NOTES,
			Text:           text,
			AppendText:     true,
			UpdateText:     true,
			Subtitle:       "[!]",
			UpdateSubtitle: false,
		},
	}

}
