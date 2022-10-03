package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"fmt"
	"time"

	"github.com/maxb-odessa/slog"
)

type FSSSignalDiscoveredEvent struct {
	IsStation                bool      `mapstructure:"IsStation,omitempty"`
	SignalName               string    `mapstructure:"SignalName,omitempty"`
	SignalNameLocalised      string    `mapstructure:"SignalName_Localised,omitempty"`
	SpawningFaction          string    `mapstructure:"SpawningFaction,omitempty"`
	SpawningFactionLocalised string    `mapstructure:"SpawningFaction_Localised,omitempty"`
	SpawningState            string    `mapstructure:"SpawningState,omitempty"`
	SpawningStateLocalised   string    `mapstructure:"SpawningState_Localised,omitempty"`
	SystemAddress            int       `mapstructure:"SystemAddress,omitempty"`
	ThreatLevel              int       `mapstructure:"ThreatLevel,omitempty"`
	TimeRemaining            float64   `mapstructure:"TimeRemaining,omitempty"`
	UssType                  string    `mapstructure:"USSType,omitempty"`
	USSTypeLocalised         string    `mapstructure:"USSType_Localised,omitempty"`
	Event                    string    `mapstructure:"event,omitempty"`
	Timestamp                time.Time `mapstructure:"timestamp,omitempty"`
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
		encEntities(name),
		encEntities(usstype),
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
