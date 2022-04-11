package file

import "time"

type OutfittingEvent struct {
	MarketID    int       `mapstructure:"MarketID,omitempty"`
	StarSystem  string    `mapstructure:"StarSystem,omitempty"`
	StationName string    `mapstructure:"StationName,omitempty"`
	Event       string    `mapstructure:"event,omitempty"`
	Timestamp   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evOutfitting(ev *OutfittingEvent) {
	return
}
