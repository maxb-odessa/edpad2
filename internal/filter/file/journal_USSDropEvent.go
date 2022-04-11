package file

import "time"

type USSDropEvent struct {
	UssThreat        int       `mapstructure:"USSThreat,omitempty"`
	UssType          string    `mapstructure:"USSType,omitempty"`
	USSTypeLocalised string    `mapstructure:"USSType_Localised,omitempty"`
	Event            string    `mapstructure:"event,omitempty"`
	Timestamp        time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evUSSDrop(ev *USSDropEvent) {
	return
}
