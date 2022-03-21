package file

import "time"

type LaunchSRVEvent struct {
	ID               int       `json:"ID,omitempty"`
	Loadout          string    `json:"Loadout,omitempty"`
	PlayerControlled bool      `json:"PlayerControlled,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLaunchSRV(ev *LaunchSRVEvent) {
	return
}
