package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"time"
)

type FuelScoopEvent struct {
	Scooped   float64   `json:"Scooped,omitempty"`
	Total     float64   `json:"Total,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFuelScoop(ev *FuelScoopEvent) {

	if FuelLevel < 12.0 && ev.Total >= 12.0 {
		h.connector.ToRouterCh <- &router.Message{
			Dst: router.LocalDisplay,
			Data: &display.Text{
				ViewPort:       display.VP_INFO,
				Text:           "",
				AppendText:     false,
				UpdateText:     true,
				Subtitle:       "",
				UpdateSubtitle: false,
			},
		}
	}

	FuelLevel = ev.Total

	return
}
