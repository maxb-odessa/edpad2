package file

import (
	"time"
)

type FSSBodySignalsEvent struct {
	BodyID   int    `json:"BodyID,omitempty"`
	BodyName string `json:"BodyName,omitempty"`
	Signals  []struct {
		Count         int    `json:"Count,omitempty"`
		Type          string `json:"Type,omitempty"`
		TypeLocalised string `json:"Type_Localised,omitempty"`
	} `json:"Signals,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
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
			refresh = true
		case "$SAA_SignalType_Biological;":
			pd.signals.biological += s.Count
		}

		refresh = true

	}

	// refresh planets table herea
	if refresh {
		h.refreshPlanets()
	}

	return
}
