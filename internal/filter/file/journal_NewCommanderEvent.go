package file

import "time"

type NewCommanderEvent struct {
	Fid       string    `json:"FID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Package   string    `json:"Package,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evNewCommander(ev *NewCommanderEvent) {
	return
}
