package file

import "time"

type PowerplaySalaryEvent struct {
	Amount    int       `json:"Amount,omitempty"`
	Power     string    `json:"Power,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPowerplaySalary(ev *PowerplaySalaryEvent) {
	return
}
