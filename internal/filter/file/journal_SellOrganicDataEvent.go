package file

import "time"

type SellOrganicDataEvent struct {
	BioData []struct {
		Bonus            int    `json:"Bonus,omitempty"`
		Genus            string `json:"Genus,omitempty"`
		GenusLocalised   string `json:"Genus_Localised,omitempty"`
		Species          string `json:"Species,omitempty"`
		SpeciesLocalised string `json:"Species_Localised,omitempty"`
		Value            int    `json:"Value,omitempty"`
	} `json:"BioData,omitempty"`
	MarketID  int       `json:"MarketID,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSellOrganicData(ev *SellOrganicDataEvent) {
	return
}
