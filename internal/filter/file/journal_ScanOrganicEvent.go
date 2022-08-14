package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"time"
)

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

	text := "Biological discovery:" + ev.GenusLocalised + ", " + ev.SpeciesLocalised + "\n"

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_NOTES,
			Text:           text,
			AppendText:     true,
			UpdateText:     true,
			Subtitle:       "[!]",
			UpdateSubtitle: true,
		},
	}

	return
}
