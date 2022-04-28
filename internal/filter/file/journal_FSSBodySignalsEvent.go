package file

import (
	"time"
)

type FSSBodySignalsEvent struct {
	BodyID   int    `mapstructure:"BodyID,omitempty"`
	BodyName string `mapstructure:"BodyName,omitempty"`
	Signals  []struct {
		Count         int    `mapstructure:"Count,omitempty"`
		Type          string `mapstructure:"Type,omitempty"`
		TypeLocalised string `mapstructure:"Type_Localised,omitempty"`
	} `mapstructure:"Signals,omitempty"`
	SystemAddress int       `mapstructure:"SystemAddress,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

/*
{ "timestamp":"2022-02-12T17:55:43Z", "event":"FSSBodySignals", "BodyName":"Umbaits ER-U d3-10 B 7", "BodyID":25, "SystemAddress":357664970019, "Signals":[ { "Type":"$SAA_SignalType_Biological;", "Type_Localised":"Biological", "Count":1 }, { "Type":"$SAA_SignalType_Geological;", "Type_Localised":"Geological", "Count":2 } ] }
*/
func (h *handler) evFSSBodySignals(ev *FSSBodySignalsEvent) {

	pd, ok := CurrentSystemPlanets[ev.BodyName]
	if !ok {
		pd = new(planetData)
		pd.signals = new(bodySignals)
		CurrentSystemPlanets[ev.BodyName] = pd
	}

	refresh := false

	for _, s := range ev.Signals {

		switch s.Type {
		case "$SAA_SignalType_Geological;":
			pd.signals.geological += s.Count
		case "$SAA_SignalType_Biological;":
			pd.signals.biological += s.Count
		case "$SAA_SignalType_Guardian;":
			pd.signals.guardian += s.Count
		case "$SAA_SignalType_Human;":
			pd.signals.human += s.Count
		case "$SAA_SignalType_Other;":
			pd.signals.other += s.Count
		}

		refresh = true

	}

	// refresh planets table here
	if refresh {
		h.refreshPlanets()
	}

	return

}
