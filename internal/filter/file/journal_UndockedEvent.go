package file

import "time"

type UndockedEvent struct {
	MarketID    int       `mapstructure:"MarketID,omitempty"`
	Multicrew   bool      `mapstructure:"Multicrew,omitempty"`
	StationName string    `mapstructure:"StationName,omitempty"`
	StationType string    `mapstructure:"StationType,omitempty"`
	Taxi        bool      `mapstructure:"Taxi,omitempty"`
	Event       string    `mapstructure:"event,omitempty"`
	Timestamp   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evUndocked(ev *UndockedEvent) {
	return
}
