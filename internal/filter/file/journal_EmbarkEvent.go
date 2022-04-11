package file

import "time"

type EmbarkEvent struct {
	Body          string    `mapstructure:"Body,omitempty"`
	BodyID        int       `mapstructure:"BodyID,omitempty"`
	ID            int       `mapstructure:"ID,omitempty"`
	MarketID      int       `mapstructure:"MarketID,omitempty"`
	Multicrew     bool      `mapstructure:"Multicrew,omitempty"`
	OnPlanet      bool      `mapstructure:"OnPlanet,omitempty"`
	OnStation     bool      `mapstructure:"OnStation,omitempty"`
	Srv           bool      `mapstructure:"SRV,omitempty"`
	StarSystem    string    `mapstructure:"StarSystem,omitempty"`
	StationName   string    `mapstructure:"StationName,omitempty"`
	StationType   string    `mapstructure:"StationType,omitempty"`
	SystemAddress int       `mapstructure:"SystemAddress,omitempty"`
	Taxi          bool      `mapstructure:"Taxi,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evEmbark(ev *EmbarkEvent) {
	return
}
