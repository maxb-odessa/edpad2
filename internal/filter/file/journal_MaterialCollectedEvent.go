package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"fmt"
	"time"
)

type MaterialCollectedEvent struct {
	Category      string    `mapstructure:"Category,omitempty"`
	Count         int       `mapstructure:"Count,omitempty"`
	Name          string    `mapstructure:"Name,omitempty"`
	NameLocalised string    `mapstructure:"Name_Localised,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
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
