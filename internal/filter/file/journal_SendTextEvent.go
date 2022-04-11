package file

import "time"

type SendTextEvent struct {
	Message   string    `mapstructure:"Message,omitempty"`
	Sent      bool      `mapstructure:"Sent,omitempty"`
	To        string    `mapstructure:"To,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evSendText(ev *SendTextEvent) {
	return
}
