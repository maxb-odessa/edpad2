package file

import "time"

type TechnologyBrokerEvent struct {
	BrokerType  string `json:"BrokerType,omitempty"`
	Commodities []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"Commodities,omitempty"`
	ItemsUnlocked []struct {
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"ItemsUnlocked,omitempty"`
	MarketID  int `json:"MarketID,omitempty"`
	Materials []struct {
		Category      string `json:"Category,omitempty"`
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"Materials,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evTechnologyBroker(ev *TechnologyBrokerEvent) {
	return
}
