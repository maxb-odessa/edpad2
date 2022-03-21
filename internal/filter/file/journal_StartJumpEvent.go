package file

import "time"

type StartJumpEvent struct {
	JumpType      string    `json:"JumpType,omitempty"`
	StarClass     string    `json:"StarClass,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evStartJump(ev *StartJumpEvent) {
	return
}
