package file

import "time"

type ApproachSettlementEvent struct {
	BodyID        int       `mapstructure:"BodyID,omitempty"`
	BodyName      string    `mapstructure:"BodyName,omitempty"`
	Latitude      float64   `mapstructure:"Latitude,omitempty"`
	Longitude     float64   `mapstructure:"Longitude,omitempty"`
	MarketID      int       `mapstructure:"MarketID,omitempty"`
	Name          string    `mapstructure:"Name,omitempty"`
	NameLocalised string    `mapstructure:"Name_Localised,omitempty"`
	SystemAddress int       `mapstructure:"SystemAddress,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evApproachSettlement(ev *ApproachSettlementEvent) {
	return
}
