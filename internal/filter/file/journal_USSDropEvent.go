package file

import "time"

type USSDropEvent struct {
	UssThreat        int       `json:"USSThreat,omitempty"`
	UssType          string    `json:"USSType,omitempty"`
	USSTypeLocalised string    `json:"USSType_Localised,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evUSSDrop(ev *USSDropEvent) {
	return
}
