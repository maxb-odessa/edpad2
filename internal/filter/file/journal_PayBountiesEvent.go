package file

import "time"

type PayBountiesEvent struct {
	Amount           int       `mapstructure:"Amount,omitempty"`
	BrokerPercentage float64   `mapstructure:"BrokerPercentage,omitempty"`
	Faction          string    `mapstructure:"Faction,omitempty"`
	FactionLocalised string    `mapstructure:"Faction_Localised,omitempty"`
	ShipID           int       `mapstructure:"ShipID,omitempty"`
	Event            string    `mapstructure:"event,omitempty"`
	Timestamp        time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evPayBounties(ev *PayBountiesEvent) {
	return
}
