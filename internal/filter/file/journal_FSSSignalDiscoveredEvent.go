package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"fmt"
	"time"

	"github.com/maxb-odessa/slog"
)

type FSSSignalDiscoveredEvent struct {
	IsStation                bool      `json:"IsStation,omitempty"`
	SignalName               string    `json:"SignalName,omitempty"`
	SignalNameLocalised      string    `json:"SignalName_Localised,omitempty"`
	SpawningFaction          string    `json:"SpawningFaction,omitempty"`
	SpawningFactionLocalised string    `json:"SpawningFaction_Localised,omitempty"`
	SpawningState            string    `json:"SpawningState,omitempty"`
	SpawningStateLocalised   string    `json:"SpawningState_Localised,omitempty"`
	SystemAddress            int       `json:"SystemAddress,omitempty"`
	ThreatLevel              int       `json:"ThreatLevel,omitempty"`
	TimeRemaining            float64   `json:"TimeRemaining,omitempty"`
	UssType                  string    `json:"USSType,omitempty"`
	USSTypeLocalised         string    `json:"USSType_Localised,omitempty"`
	Event                    string    `json:"event,omitempty"`
	Timestamp                time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFSSSignalDiscovered(ev *FSSSignalDiscoveredEvent) {

	// TODO here

	// simply print what we've got

	name := "unknown"
	if ev.SignalNameLocalised != "" {
		name = ev.SignalNameLocalised
	} else if ev.SignalName != "" {
		name = ev.SignalName
	}

	usstype := "unknown"
	if ev.USSTypeLocalised != "" {
		usstype = ev.USSTypeLocalised
	} else if ev.UssType != "" {
		usstype = ev.UssType
	}

	text := fmt.Sprintf("Name: %s\n --> Type: %s, Threat: %d, Station: %v\n",
		name,
		usstype,
		ev.ThreatLevel,
		ev.IsStation,
	)

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_SIGNALS,
			Text:           text,
			AppendText:     true,
			UpdateText:     true,
			Subtitle:       `[!]`,
			UpdateSubtitle: true,
		},
	}

	slog.Debug(5, "SIGNAL: %s\n%+v", text, ev)

	return
}
