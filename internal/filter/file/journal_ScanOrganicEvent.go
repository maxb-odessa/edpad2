package file

import "time"

type ScanOrganicEvent struct {
	Body             int       `json:"Body,omitempty"`
	Genus            string    `json:"Genus,omitempty"`
	GenusLocalised   string    `json:"Genus_Localised,omitempty"`
	ScanType         string    `json:"ScanType,omitempty"`
	Species          string    `json:"Species,omitempty"`
	SpeciesLocalised string    `json:"Species_Localised,omitempty"`
	SystemAddress    int       `json:"SystemAddress,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evScanOrganic(ev *ScanOrganicEvent) {
	return
}
