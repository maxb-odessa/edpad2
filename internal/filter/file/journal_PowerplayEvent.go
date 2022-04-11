package file

import "time"

type PowerplayEvent struct {
	Merits      int       `mapstructure:"Merits,omitempty"`
	Power       string    `mapstructure:"Power,omitempty"`
	Rank        int       `mapstructure:"Rank,omitempty"`
	TimePledged int       `mapstructure:"TimePledged,omitempty"`
	Votes       int       `mapstructure:"Votes,omitempty"`
	Event       string    `mapstructure:"event,omitempty"`
	Timestamp   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evPowerplay(ev *PowerplayEvent) {
	return
}
