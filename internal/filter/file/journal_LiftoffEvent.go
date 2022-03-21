package file

import "time"

type LiftoffEvent struct {
	Body                        string    `json:"Body,omitempty"`
	BodyID                      int       `json:"BodyID,omitempty"`
	Latitude                    float64   `json:"Latitude,omitempty"`
	Longitude                   float64   `json:"Longitude,omitempty"`
	Multicrew                   bool      `json:"Multicrew,omitempty"`
	NearestDestination          string    `json:"NearestDestination,omitempty"`
	NearestDestinationLocalised string    `json:"NearestDestination_Localised,omitempty"`
	OnPlanet                    bool      `json:"OnPlanet,omitempty"`
	OnStation                   bool      `json:"OnStation,omitempty"`
	PlayerControlled            bool      `json:"PlayerControlled,omitempty"`
	StarSystem                  string    `json:"StarSystem,omitempty"`
	SystemAddress               int       `json:"SystemAddress,omitempty"`
	Taxi                        bool      `json:"Taxi,omitempty"`
	Event                       string    `json:"event,omitempty"`
	Timestamp                   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLiftoff(ev *LiftoffEvent) {
	return
}
