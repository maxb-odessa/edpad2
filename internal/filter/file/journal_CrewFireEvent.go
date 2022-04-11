package file

import "time"

type CrewFireEvent struct {
	CrewID    int       `mapstructure:"CrewID,omitempty"`
	Name      string    `mapstructure:"Name,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCrewFire(ev *CrewFireEvent) {
	return
}
