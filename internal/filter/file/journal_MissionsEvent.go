package file

import "time"

type MissionsEvent struct {
	Active []struct {
		Expires          int    `mapstructure:"Expires,omitempty"`
		MissionID        int    `mapstructure:"MissionID,omitempty"`
		Name             string `mapstructure:"Name,omitempty"`
		PassengerMission bool   `mapstructure:"PassengerMission,omitempty"`
	} `mapstructure:"Active,omitempty"`
	Complete []interface{} `mapstructure:"Complete,omitempty"`
	Failed   []struct {
		Expires          int    `mapstructure:"Expires,omitempty"`
		MissionID        int    `mapstructure:"MissionID,omitempty"`
		Name             string `mapstructure:"Name,omitempty"`
		PassengerMission bool   `mapstructure:"PassengerMission,omitempty"`
	} `mapstructure:"Failed,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMissions(ev *MissionsEvent) {
	return
}
