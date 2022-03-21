package file

import "time"

type CrewAssignEvent struct {
	CrewID    int       `json:"CrewID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Role      string    `json:"Role,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCrewAssign(ev *CrewAssignEvent) {
	return
}
