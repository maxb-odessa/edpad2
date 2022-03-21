package file

import "time"

type FSSAllBodiesFoundEvent struct {
	Count         int       `json:"Count,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	SystemName    string    `json:"SystemName,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFSSAllBodiesFound(ev *FSSAllBodiesFoundEvent) {
	return
}
