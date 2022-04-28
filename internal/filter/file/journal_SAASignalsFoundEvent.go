package file

import "time"

type SAASignalsFoundEvent struct {
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
{ "timestamp":"2021-03-16T14:11:06Z", "event":"SAASignalsFound", "BodyName":"Synuefe NL-N c23-4 B 3", "SystemAddress":1184840454858, "BodyID":18, "Signals":[ { "Type":"$SAA_SignalType_Guardian;", "Type_Localised":"Guardian", "Count":4 }, { "Type":"$SAA_SignalType_Human;", "Type_Localised":"Human", "Count":16 } ] }
{ "timestamp":"2021-02-07T14:59:45Z", "event":"SAASignalsFound", "BodyName":"Moon", "SystemAddress":10477373803, "BodyID":4, "Signals":[ { "Type":"$SAA_SignalType_Other;", "Type_Localised":"Other", "Count":1 }, { "Type":"$SAA_SignalType_Human;", "Type_Localised":"Human", "Count":1 } ] }
*/

// Surface Scanner, not FSS !
func (h *handler) evSAASignalsFound(ev *SAASignalsFoundEvent) {

	/*
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
	*/
	return
}
