package file

import "time"

type MissionsEvent struct {
	Active []struct {
		Expires          int    `json:"Expires,omitempty"`
		MissionID        int    `json:"MissionID,omitempty"`
		Name             string `json:"Name,omitempty"`
		PassengerMission bool   `json:"PassengerMission,omitempty"`
	} `json:"Active,omitempty"`
	Complete []interface{} `json:"Complete,omitempty"`
	Failed   []struct {
		Expires          int    `json:"Expires,omitempty"`
		MissionID        int    `json:"MissionID,omitempty"`
		Name             string `json:"Name,omitempty"`
		PassengerMission bool   `json:"PassengerMission,omitempty"`
	} `json:"Failed,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMissions(ev *MissionsEvent) {
	return
}
