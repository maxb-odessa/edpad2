package file

import "time"

type ReceiveTextEvent struct {
	Channel          string    `mapstructure:"Channel,omitempty"`
	From             string    `mapstructure:"From,omitempty"`
	FromLocalised    string    `mapstructure:"From_Localised,omitempty"`
	Message          string    `mapstructure:"Message,omitempty"`
	MessageLocalised string    `mapstructure:"Message_Localised,omitempty"`
	Event            string    `mapstructure:"event,omitempty"`
	Timestamp        time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evReceiveText(ev *ReceiveTextEvent) {
	return
}
