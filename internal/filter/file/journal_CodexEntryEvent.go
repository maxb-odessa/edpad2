package file

import "time"

type CodexEntryEvent struct {
	Category                    string    `json:"Category,omitempty"`
	CategoryLocalised           string    `json:"Category_Localised,omitempty"`
	EntryID                     int       `json:"EntryID,omitempty"`
	IsNewEntry                  bool      `json:"IsNewEntry,omitempty"`
	Latitude                    float64   `json:"Latitude,omitempty"`
	Longitude                   float64   `json:"Longitude,omitempty"`
	Name                        string    `json:"Name,omitempty"`
	NameLocalised               string    `json:"Name_Localised,omitempty"`
	NearestDestination          string    `json:"NearestDestination,omitempty"`
	NearestDestinationLocalised string    `json:"NearestDestination_Localised,omitempty"`
	Region                      string    `json:"Region,omitempty"`
	RegionLocalised             string    `json:"Region_Localised,omitempty"`
	SubCategory                 string    `json:"SubCategory,omitempty"`
	SubCategoryLocalised        string    `json:"SubCategory_Localised,omitempty"`
	System                      string    `json:"System,omitempty"`
	SystemAddress               int       `json:"SystemAddress,omitempty"`
	VoucherAmount               int       `json:"VoucherAmount,omitempty"`
	Event                       string    `json:"event,omitempty"`
	Timestamp                   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCodexEntry(ev *CodexEntryEvent) {
	return
}
