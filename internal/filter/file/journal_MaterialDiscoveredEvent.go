package file

import "time"

type MaterialDiscoveredEvent struct {
	Category        string    `mapstructure:"Category,omitempty"`
	DiscoveryNumber int       `mapstructure:"DiscoveryNumber,omitempty"`
	Name            string    `mapstructure:"Name,omitempty"`
	NameLocalised   string    `mapstructure:"Name_Localised,omitempty"`
	Event           string    `mapstructure:"event,omitempty"`
	Timestamp       time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMaterialDiscovered(ev *MaterialDiscoveredEvent) {
	return
}
