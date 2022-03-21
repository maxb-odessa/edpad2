package file

import "time"

type SendTextEvent struct {
	Message   string    `json:"Message,omitempty"`
	Sent      bool      `json:"Sent,omitempty"`
	To        string    `json:"To,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSendText(ev *SendTextEvent) {
	return
}
