package file

import "time"

type ResurrectEvent struct {
	Bankrupt  bool      `mapstructure:"Bankrupt,omitempty"`
	Cost      int       `mapstructure:"Cost,omitempty"`
	Option    string    `mapstructure:"Option,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evResurrect(ev *ResurrectEvent) {
	return
}
