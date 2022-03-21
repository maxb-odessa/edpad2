package file

import "time"

type ReceiveTextEvent struct {
	Channel          string    `json:"Channel,omitempty"`
	From             string    `json:"From,omitempty"`
	FromLocalised    string    `json:"From_Localised,omitempty"`
	Message          string    `json:"Message,omitempty"`
	MessageLocalised string    `json:"Message_Localised,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evReceiveText(ev *ReceiveTextEvent) {
	return
}
