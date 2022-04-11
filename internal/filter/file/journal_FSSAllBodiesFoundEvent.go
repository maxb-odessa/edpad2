package file

import "time"

type FSSAllBodiesFoundEvent struct {
	Count         int       `mapstructure:"Count,omitempty"`
	SystemAddress int       `mapstructure:"SystemAddress,omitempty"`
	SystemName    string    `mapstructure:"SystemName,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evFSSAllBodiesFound(ev *FSSAllBodiesFoundEvent) {
	return
}
