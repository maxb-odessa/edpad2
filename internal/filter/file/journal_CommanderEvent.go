package file

import "time"

type CommanderEvent struct {
	Fid       string    `json:"FID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCommander(ev *CommanderEvent) {
	return
}
