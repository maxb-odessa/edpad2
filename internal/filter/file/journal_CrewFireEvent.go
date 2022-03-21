package file

import "time"

type CrewFireEvent struct {
	CrewID    int       `json:"CrewID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCrewFire(ev *CrewFireEvent) {
	return
}
