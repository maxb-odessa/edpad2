package file

import "time"

type PowerplaySalaryEvent struct {
	Amount    int       `mapstructure:"Amount,omitempty"`
	Power     string    `mapstructure:"Power,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evPowerplaySalary(ev *PowerplaySalaryEvent) {
	return
}
