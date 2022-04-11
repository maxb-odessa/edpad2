package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"time"
)

type FuelScoopEvent struct {
	Scooped   float64   `mapstructure:"Scooped,omitempty"`
	Total     float64   `mapstructure:"Total,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
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
