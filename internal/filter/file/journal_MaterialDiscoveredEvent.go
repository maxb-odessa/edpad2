package file

import "time"

type MaterialDiscoveredEvent struct {
	Category        string    `json:"Category,omitempty"`
	DiscoveryNumber int       `json:"DiscoveryNumber,omitempty"`
	Name            string    `json:"Name,omitempty"`
	NameLocalised   string    `json:"Name_Localised,omitempty"`
	Event           string    `json:"event,omitempty"`
	Timestamp       time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMaterialDiscovered(ev *MaterialDiscoveredEvent) {
	return
}
