package file

import "time"

type RebootRepairEvent struct {
	Modules   []interface{} `json:"Modules,omitempty"`
	Event     string        `json:"event,omitempty"`
	Timestamp time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evRebootRepair(ev *RebootRepairEvent) {
	return
}
