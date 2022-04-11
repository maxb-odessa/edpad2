package file

import "time"

type MissionRedirectedEvent struct {
	MissionID             int       `mapstructure:"MissionID,omitempty"`
	Name                  string    `mapstructure:"Name,omitempty"`
	NewDestinationStation string    `mapstructure:"NewDestinationStation,omitempty"`
	NewDestinationSystem  string    `mapstructure:"NewDestinationSystem,omitempty"`
	OldDestinationStation string    `mapstructure:"OldDestinationStation,omitempty"`
	OldDestinationSystem  string    `mapstructure:"OldDestinationSystem,omitempty"`
	Event                 string    `mapstructure:"event,omitempty"`
	Timestamp             time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMissionRedirected(ev *MissionRedirectedEvent) {
	return
}
