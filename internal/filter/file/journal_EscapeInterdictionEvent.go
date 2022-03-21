package file

import "time"

type EscapeInterdictionEvent struct {
	Interdictor string    `json:"Interdictor,omitempty"`
	IsPlayer    bool      `json:"IsPlayer,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evEscapeInterdiction(ev *EscapeInterdictionEvent) {
	return
}
