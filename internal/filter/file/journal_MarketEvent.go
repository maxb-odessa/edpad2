package file

import "time"

type MarketEvent struct {
	CarrierDockingAccess string    `mapstructure:"CarrierDockingAccess,omitempty"`
	MarketID             int       `mapstructure:"MarketID,omitempty"`
	StarSystem           string    `mapstructure:"StarSystem,omitempty"`
	StationName          string    `mapstructure:"StationName,omitempty"`
	StationType          string    `mapstructure:"StationType,omitempty"`
	Event                string    `mapstructure:"event,omitempty"`
	Timestamp            time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMarket(ev *MarketEvent) {
	return
}
