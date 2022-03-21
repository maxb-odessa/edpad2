package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"encoding/json"
	"fmt"
	"time"
)

type ScanEvent struct {
	AbsoluteMagnitude     float64 `json:"AbsoluteMagnitude,omitempty"`
	AgeMy                 int     `json:"Age_MY,omitempty"`
	Atmosphere            string  `json:"Atmosphere,omitempty"`
	AtmosphereComposition []struct {
		Name    string  `json:"Name,omitempty"`
		Percent float64 `json:"Percent,omitempty"`
	} `json:"AtmosphereComposition,omitempty"`
	AtmosphereType string  `json:"AtmosphereType,omitempty"`
	AxialTilt      float64 `json:"AxialTilt,omitempty"`
	BodyID         int     `json:"BodyID,omitempty"`
	BodyName       string  `json:"BodyName,omitempty"`
	Composition    *struct {
		Ice   float64 `json:"Ice,omitempty"`
		Metal float64 `json:"Metal,omitempty"`
		Rock  float64 `json:"Rock,omitempty"`
	} `json:"Composition,omitempty"`
	DistanceFromArrivalLs float64 `json:"DistanceFromArrivalLS,omitempty"`
	Eccentricity          float64 `json:"Eccentricity,omitempty"`
	Landable              bool    `json:"Landable,omitempty"`
	Luminosity            string  `json:"Luminosity,omitempty"`
	MassEm                float64 `json:"MassEM,omitempty"`
	Materials             []struct {
		Name    string  `json:"Name,omitempty"`
		Percent float64 `json:"Percent,omitempty"`
	} `json:"Materials,omitempty"`
	OrbitalInclination float64 `json:"OrbitalInclination,omitempty"`
	OrbitalPeriod      float64 `json:"OrbitalPeriod,omitempty"`
	Parents            []struct {
		Null   int `json:"Null,omitempty"`
		Planet int `json:"Planet,omitempty"`
		Ring   int `json:"Ring,omitempty"`
		Star   int `json:"Star,omitempty"`
	} `json:"Parents,omitempty"`
	Periapsis    float64 `json:"Periapsis,omitempty"`
	PlanetClass  string  `json:"PlanetClass,omitempty"`
	Radius       float64 `json:"Radius,omitempty"`
	ReserveLevel string  `json:"ReserveLevel,omitempty"`
	Rings        []struct {
		InnerRad  json.Number `json:"InnerRad,omitempty"`
		MassMt    json.Number `json:"MassMT,omitempty"`
		Name      string      `json:"Name,omitempty"`
		OuterRad  float64     `json:"OuterRad,omitempty"`
		RingClass string      `json:"RingClass,omitempty"`
	} `json:"Rings,omitempty"`
	RotationPeriod     float64   `json:"RotationPeriod,omitempty"`
	ScanType           string    `json:"ScanType,omitempty"`
	SemiMajorAxis      float64   `json:"SemiMajorAxis,omitempty"`
	StarSystem         string    `json:"StarSystem,omitempty"`
	StarType           string    `json:"StarType,omitempty"`
	StellarMass        float64   `json:"StellarMass,omitempty"`
	Subclass           int       `json:"Subclass,omitempty"`
	SurfaceGravity     float64   `json:"SurfaceGravity,omitempty"`
	SurfacePressure    float64   `json:"SurfacePressure,omitempty"`
	SurfaceTemperature float64   `json:"SurfaceTemperature,omitempty"`
	SystemAddress      int       `json:"SystemAddress,omitempty"`
	TerraformState     string    `json:"TerraformState,omitempty"`
	TidalLock          bool      `json:"TidalLock,omitempty"`
	Volcanism          string    `json:"Volcanism,omitempty"`
	WasDiscovered      bool      `json:"WasDiscovered,omitempty"`
	WasMapped          bool      `json:"WasMapped,omitempty"`
	Event              string    `json:"event,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evScan(ev *ScanEvent) {
	m := &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_SRV,
			Text:           fmt.Sprintf("%+v", ev),
			AppendText:     true,
			UpdateText:     true,
			Subtitle:       "(123)",
			UpdateSubtitle: true,
		},
	}
	h.connector.ToRouterCh <- m
	return
}
