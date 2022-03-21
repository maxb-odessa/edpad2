package file

import "time"

type CargoDepotEvent struct {
	CargoType           string    `json:"CargoType,omitempty"`
	CargoTypeLocalised  string    `json:"CargoType_Localised,omitempty"`
	Count               int       `json:"Count,omitempty"`
	EndMarketID         int       `json:"EndMarketID,omitempty"`
	ItemsCollected      int       `json:"ItemsCollected,omitempty"`
	ItemsDelivered      int       `json:"ItemsDelivered,omitempty"`
	MissionID           int       `json:"MissionID,omitempty"`
	Progress            float64   `json:"Progress,omitempty"`
	StartMarketID       int       `json:"StartMarketID,omitempty"`
	TotalItemsToDeliver int       `json:"TotalItemsToDeliver,omitempty"`
	UpdateType          string    `json:"UpdateType,omitempty"`
	Event               string    `json:"event,omitempty"`
	Timestamp           time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCargoDepot(ev *CargoDepotEvent) {
	return
}
