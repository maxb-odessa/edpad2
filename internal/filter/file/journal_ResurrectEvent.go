package file

import "time"

type ResurrectEvent struct {
	Bankrupt  bool      `json:"Bankrupt,omitempty"`
	Cost      int       `json:"Cost,omitempty"`
	Option    string    `json:"Option,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evResurrect(ev *ResurrectEvent) {
	return
}
