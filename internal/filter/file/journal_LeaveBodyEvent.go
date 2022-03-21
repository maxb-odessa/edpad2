package file

import "time"

type LeaveBodyEvent struct {
	Body          string    `json:"Body,omitempty"`
	BodyID        int       `json:"BodyID,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLeaveBody(ev *LeaveBodyEvent) {
	return
}
