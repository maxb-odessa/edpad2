package file

import "time"

type MissionFailedEvent struct {
	MissionID int       `json:"MissionID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMissionFailed(ev *MissionFailedEvent) {
	return
}
