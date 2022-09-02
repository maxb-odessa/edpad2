package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"time"
)

type CodexEntryEvent struct {
	Category                    string    `mapstructure:"Category,omitempty"`
	CategoryLocalised           string    `mapstructure:"Category_Localised,omitempty"`
	EntryID                     int       `mapstructure:"EntryID,omitempty"`
	IsNewEntry                  bool      `mapstructure:"IsNewEntry,omitempty"`
	Latitude                    float64   `mapstructure:"Latitude,omitempty"`
	Longitude                   float64   `mapstructure:"Longitude,omitempty"`
	Name                        string    `mapstructure:"Name,omitempty"`
	NameLocalised               string    `mapstructure:"Name_Localised,omitempty"`
	NearestDestination          string    `mapstructure:"NearestDestination,omitempty"`
	NearestDestinationLocalised string    `mapstructure:"NearestDestination_Localised,omitempty"`
	Region                      string    `mapstructure:"Region,omitempty"`
	RegionLocalised             string    `mapstructure:"Region_Localised,omitempty"`
	SubCategory                 string    `mapstructure:"SubCategory,omitempty"`
	SubCategoryLocalised        string    `mapstructure:"SubCategory_Localised,omitempty"`
	System                      string    `mapstructure:"System,omitempty"`
	SystemAddress               int       `mapstructure:"SystemAddress,omitempty"`
	VoucherAmount               int       `mapstructure:"VoucherAmount,omitempty"`
	Event                       string    `mapstructure:"event,omitempty"`
	Timestamp                   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCodexEntry(ev *CodexEntryEvent) {

	isNew := ""
	if !ev.IsNewEntry {
		isNew = " (NEW)"
	}

	text := "Codex" + isNew + ":" + ev.CategoryLocalised + ", " + ev.SubCategoryLocalised + "\n" + `  +- ` +
		ev.NameLocalised + ", region: " + ev.RegionLocalised + "\n"

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_LOGS,
			Text:           text,
			AppendText:     true,
			UpdateText:     true,
			Subtitle:       "[!]",
			UpdateSubtitle: true,
		},
	}

}
