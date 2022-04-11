package file

import "time"

type NewCommanderEvent struct {
	Fid       string    `mapstructure:"FID,omitempty"`
	Name      string    `mapstructure:"Name,omitempty"`
	Package   string    `mapstructure:"Package,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evNewCommander(ev *NewCommanderEvent) {
	return
}
