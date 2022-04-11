package file

import "time"

type CargoDepotEvent struct {
	CargoType           string    `mapstructure:"CargoType,omitempty"`
	CargoTypeLocalised  string    `mapstructure:"CargoType_Localised,omitempty"`
	Count               int       `mapstructure:"Count,omitempty"`
	EndMarketID         int       `mapstructure:"EndMarketID,omitempty"`
	ItemsCollected      int       `mapstructure:"ItemsCollected,omitempty"`
	ItemsDelivered      int       `mapstructure:"ItemsDelivered,omitempty"`
	MissionID           int       `mapstructure:"MissionID,omitempty"`
	Progress            float64   `mapstructure:"Progress,omitempty"`
	StartMarketID       int       `mapstructure:"StartMarketID,omitempty"`
	TotalItemsToDeliver int       `mapstructure:"TotalItemsToDeliver,omitempty"`
	UpdateType          string    `mapstructure:"UpdateType,omitempty"`
	Event               string    `mapstructure:"event,omitempty"`
	Timestamp           time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCargoDepot(ev *CargoDepotEvent) {
	return
}
