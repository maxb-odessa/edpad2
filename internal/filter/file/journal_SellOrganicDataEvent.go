package file

import "time"

type SellOrganicDataEvent struct {
	BioData []struct {
		Bonus            int    `mapstructure:"Bonus,omitempty"`
		Genus            string `mapstructure:"Genus,omitempty"`
		GenusLocalised   string `mapstructure:"Genus_Localised,omitempty"`
		Species          string `mapstructure:"Species,omitempty"`
		SpeciesLocalised string `mapstructure:"Species_Localised,omitempty"`
		Value            int    `mapstructure:"Value,omitempty"`
	} `mapstructure:"BioData,omitempty"`
	MarketID  int       `mapstructure:"MarketID,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evSellOrganicData(ev *SellOrganicDataEvent) {
	return
}
