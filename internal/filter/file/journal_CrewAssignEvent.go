package file

import "time"

type CrewAssignEvent struct {
	CrewID    int       `mapstructure:"CrewID,omitempty"`
	Name      string    `mapstructure:"Name,omitempty"`
	Role      string    `mapstructure:"Role,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCrewAssign(ev *CrewAssignEvent) {
	return
}
