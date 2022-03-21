package file

import "time"

type SAAScanCompleteEvent struct {
	BodyID           int       `json:"BodyID,omitempty"`
	BodyName         string    `json:"BodyName,omitempty"`
	EfficiencyTarget int       `json:"EfficiencyTarget,omitempty"`
	ProbesUsed       int       `json:"ProbesUsed,omitempty"`
	SystemAddress    int       `json:"SystemAddress,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSAAScanComplete(ev *SAAScanCompleteEvent) {
	return
}
