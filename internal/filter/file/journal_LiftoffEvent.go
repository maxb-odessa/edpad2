package file

import "time"

type LiftoffEvent struct {
	Body                        string    `mapstructure:"Body,omitempty"`
	BodyID                      int       `mapstructure:"BodyID,omitempty"`
	Latitude                    float64   `mapstructure:"Latitude,omitempty"`
	Longitude                   float64   `mapstructure:"Longitude,omitempty"`
	Multicrew                   bool      `mapstructure:"Multicrew,omitempty"`
	NearestDestination          string    `mapstructure:"NearestDestination,omitempty"`
	NearestDestinationLocalised string    `mapstructure:"NearestDestination_Localised,omitempty"`
	OnPlanet                    bool      `mapstructure:"OnPlanet,omitempty"`
	OnStation                   bool      `mapstructure:"OnStation,omitempty"`
	PlayerControlled            bool      `mapstructure:"PlayerControlled,omitempty"`
	StarSystem                  string    `mapstructure:"StarSystem,omitempty"`
	SystemAddress               int       `mapstructure:"SystemAddress,omitempty"`
	Taxi                        bool      `mapstructure:"Taxi,omitempty"`
	Event                       string    `mapstructure:"event,omitempty"`
	Timestamp                   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evLiftoff(ev *LiftoffEvent) {
	return
}
