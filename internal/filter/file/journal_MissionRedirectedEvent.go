package file

import "time"

type MissionRedirectedEvent struct {
	MissionID             int       `json:"MissionID,omitempty"`
	Name                  string    `json:"Name,omitempty"`
	NewDestinationStation string    `json:"NewDestinationStation,omitempty"`
	NewDestinationSystem  string    `json:"NewDestinationSystem,omitempty"`
	OldDestinationStation string    `json:"OldDestinationStation,omitempty"`
	OldDestinationSystem  string    `json:"OldDestinationSystem,omitempty"`
	Event                 string    `json:"event,omitempty"`
	Timestamp             time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMissionRedirected(ev *MissionRedirectedEvent) {
	return
}
