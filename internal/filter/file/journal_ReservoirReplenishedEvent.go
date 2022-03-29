package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"time"
)

type ReservoirReplenishedEvent struct {
	FuelMain      float64   `json:"FuelMain,omitempty"`
	FuelReservoir float64   `json:"FuelReservoir,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evReservoirReplenished(ev *ReservoirReplenishedEvent) {

	if FuelLevel < 12.0 && ev.FuelMain >= 12.0 {
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

	FuelLevel = ev.FuelMain

	return
}
