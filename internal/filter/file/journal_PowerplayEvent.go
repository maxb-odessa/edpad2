package file

import "time"

type PowerplayEvent struct {
	Merits      int       `json:"Merits,omitempty"`
	Power       string    `json:"Power,omitempty"`
	Rank        int       `json:"Rank,omitempty"`
	TimePledged int       `json:"TimePledged,omitempty"`
	Votes       int       `json:"Votes,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPowerplay(ev *PowerplayEvent) {
	return
}
