package file

import "time"

type SAAScanCompleteEvent struct {
	BodyID           int       `mapstructure:"BodyID,omitempty"`
	BodyName         string    `mapstructure:"BodyName,omitempty"`
	EfficiencyTarget int       `mapstructure:"EfficiencyTarget,omitempty"`
	ProbesUsed       int       `mapstructure:"ProbesUsed,omitempty"`
	SystemAddress    int       `mapstructure:"SystemAddress,omitempty"`
	Event            string    `mapstructure:"event,omitempty"`
	Timestamp        time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evSAAScanComplete(ev *SAAScanCompleteEvent) {
	return
}
