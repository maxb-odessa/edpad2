package file

import "time"

type MissionAbandonedEvent struct {
	MissionID int       `mapstructure:"MissionID,omitempty"`
	Name      string    `mapstructure:"Name,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMissionAbandoned(ev *MissionAbandonedEvent) {
	return
}
