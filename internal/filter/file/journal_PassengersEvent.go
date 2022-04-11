package file

import "time"

type PassengersEvent struct {
	Manifest []struct {
		Count     int    `mapstructure:"Count,omitempty"`
		MissionID int    `mapstructure:"MissionID,omitempty"`
		Type      string `mapstructure:"Type,omitempty"`
		Vip       bool   `mapstructure:"VIP,omitempty"`
		Wanted    bool   `mapstructure:"Wanted,omitempty"`
	} `mapstructure:"Manifest,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evPassengers(ev *PassengersEvent) {
	return
}
