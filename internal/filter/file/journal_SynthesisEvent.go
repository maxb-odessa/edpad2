package file

import "time"

type SynthesisEvent struct {
	Materials []struct {
		Count int    `mapstructure:"Count,omitempty"`
		Name  string `mapstructure:"Name,omitempty"`
	} `mapstructure:"Materials,omitempty"`
	Name      string    `mapstructure:"Name,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evSynthesis(ev *SynthesisEvent) {
	return
}
