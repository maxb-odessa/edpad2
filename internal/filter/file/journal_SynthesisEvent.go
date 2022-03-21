package file

import "time"

type SynthesisEvent struct {
	Materials []struct {
		Count int    `json:"Count,omitempty"`
		Name  string `json:"Name,omitempty"`
	} `json:"Materials,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSynthesis(ev *SynthesisEvent) {
	return
}
