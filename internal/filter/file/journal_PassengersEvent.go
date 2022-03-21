package file

import "time"

type PassengersEvent struct {
	Manifest []struct {
		Count     int    `json:"Count,omitempty"`
		MissionID int    `json:"MissionID,omitempty"`
		Type      string `json:"Type,omitempty"`
		Vip       bool   `json:"VIP,omitempty"`
		Wanted    bool   `json:"Wanted,omitempty"`
	} `json:"Manifest,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPassengers(ev *PassengersEvent) {
	return
}
