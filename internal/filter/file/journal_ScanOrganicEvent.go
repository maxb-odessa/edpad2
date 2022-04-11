package file

import "time"

type ScanOrganicEvent struct {
	Body             int       `mapstructure:"Body,omitempty"`
	Genus            string    `mapstructure:"Genus,omitempty"`
	GenusLocalised   string    `mapstructure:"Genus_Localised,omitempty"`
	ScanType         string    `mapstructure:"ScanType,omitempty"`
	Species          string    `mapstructure:"Species,omitempty"`
	SpeciesLocalised string    `mapstructure:"Species_Localised,omitempty"`
	SystemAddress    int       `mapstructure:"SystemAddress,omitempty"`
	Event            string    `mapstructure:"event,omitempty"`
	Timestamp        time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evScanOrganic(ev *ScanOrganicEvent) {
	return
}
