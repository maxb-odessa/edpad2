package file

import "time"

type ReservoirReplenishedEvent struct {
	FuelMain      float64   `json:"FuelMain,omitempty"`
	FuelReservoir float64   `json:"FuelReservoir,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evReservoirReplenished(ev *ReservoirReplenishedEvent) {
	return
}
