package file

import "time"

type PayBountiesEvent struct {
	Amount           int       `json:"Amount,omitempty"`
	BrokerPercentage float64   `json:"BrokerPercentage,omitempty"`
	Faction          string    `json:"Faction,omitempty"`
	FactionLocalised string    `json:"Faction_Localised,omitempty"`
	ShipID           int       `json:"ShipID,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPayBounties(ev *PayBountiesEvent) {
	return
}
