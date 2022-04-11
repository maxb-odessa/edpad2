package file

import "time"

type EscapeInterdictionEvent struct {
	Interdictor string    `mapstructure:"Interdictor,omitempty"`
	IsPlayer    bool      `mapstructure:"IsPlayer,omitempty"`
	Event       string    `mapstructure:"event,omitempty"`
	Timestamp   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evEscapeInterdiction(ev *EscapeInterdictionEvent) {
	return
}
