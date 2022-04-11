package file

import "time"

type RebootRepairEvent struct {
	Modules   []interface{} `mapstructure:"Modules,omitempty"`
	Event     string        `mapstructure:"event,omitempty"`
	Timestamp time.Time     `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evRebootRepair(ev *RebootRepairEvent) {
	return
}
