package file

import "time"

type CommanderEvent struct {
	Fid       string    `mapstructure:"FID,omitempty"`
	Name      string    `mapstructure:"Name,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCommander(ev *CommanderEvent) {
	return
}
