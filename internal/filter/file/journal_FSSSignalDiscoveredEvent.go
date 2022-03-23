package file

import "time"

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

	return
}
