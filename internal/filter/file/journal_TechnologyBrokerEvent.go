package file

import "time"

type TechnologyBrokerEvent struct {
	BrokerType  string `mapstructure:"BrokerType,omitempty"`
	Commodities []struct {
		Count         int    `mapstructure:"Count,omitempty"`
		Name          string `mapstructure:"Name,omitempty"`
		NameLocalised string `mapstructure:"Name_Localised,omitempty"`
	} `mapstructure:"Commodities,omitempty"`
	ItemsUnlocked []struct {
		Name          string `mapstructure:"Name,omitempty"`
		NameLocalised string `mapstructure:"Name_Localised,omitempty"`
	} `mapstructure:"ItemsUnlocked,omitempty"`
	MarketID  int `mapstructure:"MarketID,omitempty"`
	Materials []struct {
		Category      string `mapstructure:"Category,omitempty"`
		Count         int    `mapstructure:"Count,omitempty"`
		Name          string `mapstructure:"Name,omitempty"`
		NameLocalised string `mapstructure:"Name_Localised,omitempty"`
	} `mapstructure:"Materials,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evTechnologyBroker(ev *TechnologyBrokerEvent) {
	return
}
