package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"fmt"
	"time"
)

type FSSDiscoveryScanEvent struct {
	BodyCount     int       `mapstructure:"BodyCount,omitempty"`
	NonBodyCount  int       `mapstructure:"NonBodyCount,omitempty"`
	Progress      float64   `mapstructure:"Progress,omitempty"`
	SystemAddress int       `mapstructure:"SystemAddress,omitempty"`
	SystemName    string    `mapstructure:"SystemName,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evFSSDiscoveryScan(ev *FSSDiscoveryScanEvent) {

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_SYSTEM,
			Text:           "",
			AppendText:     false,
			UpdateText:     false,
			Subtitle:       fmt.Sprintf("[%d/%d]", ev.BodyCount, ev.NonBodyCount),
			UpdateSubtitle: true,
		},
	}

	return
}
