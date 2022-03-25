package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"fmt"
	"time"
)

type FSSDiscoveryScanEvent struct {
	BodyCount     int       `json:"BodyCount,omitempty"`
	NonBodyCount  int       `json:"NonBodyCount,omitempty"`
	Progress      float64   `json:"Progress,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	SystemName    string    `json:"SystemName,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
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
