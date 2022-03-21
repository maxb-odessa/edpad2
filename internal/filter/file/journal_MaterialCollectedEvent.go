package file

import "time"

type MaterialCollectedEvent struct {
	Category      string    `json:"Category,omitempty"`
	Count         int       `json:"Count,omitempty"`
	Name          string    `json:"Name,omitempty"`
	NameLocalised string    `json:"Name_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMaterialCollected(ev *MaterialCollectedEvent) {
	return
}
