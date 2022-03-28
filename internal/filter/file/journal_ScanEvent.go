package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"encoding/json"
	"fmt"
	"time"

	"github.com/maxb-odessa/slog"
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

	if ev.StarType != "" {
		h.parseStar(ev)
	} else if ev.PlanetClass != "" {
		h.parsePlanet(ev)
	} else {
		// ateroids? other?
	}
}

// in meters
const (
	SOLAR_RADIUS     = 696340000.0
	EARTH_RADIUS     = 6371.0 * 1000.0
	LIGHT_SECOND     = 299792.0 * 1000.
	MIN_RING_OUT_RAD = 20.0 * LIGHT_SECOND
)

func formatTemp(temp float64) string {
	if temp > 1000000.0 {
		return fmt.Sprintf("%3.1fM", temp/1000000.0)
	} else if temp > 1000 {
		return fmt.Sprintf("%3.1fK", temp/1000.0)
	} else {
		return fmt.Sprintf("%4.0f", temp)
	}
}

func (h *handler) parseStar(ev *ScanEvent) {

	var sd starData

	sd.discovered = ev.WasDiscovered
	if ev.BodyName == CurrentMainStarName {
		sd.isMain = true
	}
	sd.class = CB(ev.StarType, 4)
	sd.subClass = ev.Subclass
	sd.distance = ev.DistanceFromArrivalLs
	sd.luminosity = ev.Luminosity
	sd.massSol = ev.StellarMass
	sd.radiusSol = ev.Radius / SOLAR_RADIUS
	sd.temperatureK = formatTemp(ev.SurfaceTemperature)
	if len(ev.Rings) > 0 {
		sd.hasBelt = true
	}

	CurrentSystemStars[ev.BodyName] = sd

	text := "\n" +
		`   <i><u><span color="gray">` +
		`  Class   Distance(ls)   Disco  Belt  M(sol)  R(sol)  Temp(K)` +
		`</span></u></i>` +
		"\n"

	yes := `<span color="yellow">yes</span>`
	no := `<span color="gray">no </span>`

	idx := 0

	for _, s := range CurrentSystemStars {

		if (idx % 2) != 0 {
			text += `<span bgcolor="#202020">`
		} else {
			text += `<span>`
		}

		mainstar := " "
		if s.isMain {
			mainstar = "*"
		}

		discovered := no
		if s.discovered {
			discovered = yes
		}

		belt := no
		if s.hasBelt {
			belt = yes
		}

		text += fmt.Sprintf(" %s %-s%-1d %-3.3s     %8.0f   %s    %s   %3.1f     %3.1f     %s",
			mainstar,
			s.class,
			s.subClass,
			s.luminosity,
			s.distance,
			discovered,
			belt,
			s.massSol,
			s.radiusSol,
			s.temperatureK)

		text += "</span>\n"

		idx++
	}

	slog.Debug(99, "display system: %+v\n%s", sd, text)

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_SYSTEM,
			Text:           text,
			AppendText:     false,
			UpdateText:     true,
			Subtitle:       "",
			UpdateSubtitle: false,
		},
	}

	return
}

func (h *handler) parsePlanet(ev *ScanEvent) {

	var pd planetData
	//  calc short name from BodyName, like "Graea Hypue FV-X d1-30 2 a" > "2 a"
	pd.shortName = ev.BodyName[len(ev.BodyName)-8:]
	pd.class = ev.PlanetClass
	pd.discovered = ev.WasDiscovered
	pd.mapped = ev.WasMapped
	pd.massEm = ev.MassEm
	pd.radiusEm = ev.Radius / EARTH_RADIUS
	pd.gravityG = ev.SurfaceGravity / 10.0
	pd.temperatureK = ev.SurfaceTemperature
	pd.rings = len(ev.Rings)
	pd.landable = ev.Landable
	if ev.TerraformState != "" {
		pd.terraformable = true
	}
	// not show, for calc only
	pd.atmosphere = ev.AtmosphereType
	pd.possibleBio = "-----" // TODO calc bios

	// calc if has wide ring
	for _, r := range ev.Rings {
		if pd.ringRad < r.OuterRad {
			pd.ringRad = r.OuterRad
		}
	}

	CurrentSystemPlanets[ev.BodyName] = pd

	text := ` <i><u><span color="gray">` +
		` Name    Type  D M  M(e)  R(e)  Grav  T(K)  Rn  Rr  Ld TF Bio  ` +
		`</span></u></i>` +
		"\n"

	idx := 0

	yes := `<span color="yellow">y</span>`
	no := `<span color="gray">n</span>`

	for _, p := range CurrentSystemPlanets {

		if !remarkablePlanet(p) {
			continue
		}

		ringRad := fmt.Sprintf("%3.0f", p.ringRad/LIGHT_SECOND)
		if p.ringRad >= MIN_RING_OUT_RAD {
			ringRad = `<span color="white">` + ringRad + `</span>`
		}

		if (idx % 2) != 0 {
			text += `<span bgcolor="#202020">`
		} else {
			text += `<span>`
		}

		discovered := no
		if p.discovered {
			discovered = yes
		}

		mapped := no
		if p.mapped {
			mapped = yes
		}

		terraformable := no
		if p.terraformable {
			terraformable = yes
		}

		landable := no
		if p.landable {
			landable = yes
		}

		text += fmt.Sprintf(" %-8.8s %s %s %s %5.2f %5.2f %5.2f %s  %2d  %s  %s  %s  %5s",
			p.shortName,
			CB(p.class, -5),
			discovered,
			mapped,
			p.massEm,
			p.radiusEm,
			p.gravityG,
			formatTemp(p.temperatureK),
			p.rings,
			ringRad,
			landable,
			terraformable,
			pd.possibleBio,
		)

		text += "</span>\n"

		idx++

	}

	slog.Debug(99, "display planets: %+v\n%s", pd, text)
	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_PLANETS,
			Text:           text,
			AppendText:     false,
			UpdateText:     true,
			Subtitle:       fmt.Sprintf("[%d]", len(CurrentSystemPlanets)),
			UpdateSubtitle: true,
		},
	}

	return
}

func remarkablePlanet(pd planetData) bool {

	// landable + high G
	if pd.landable && pd.gravityG >= 2.0 {
		return true
	}

	// many rings
	if pd.rings >= 5 {
		return true
	}

	// wide ring
	if pd.ringRad > MIN_RING_OUT_RAD {
		return true
	}

	// terraformable
	if pd.terraformable {
		return true
	}

	// heliums ar nice
	if pd.class[0:6] == "Helium" {
		return true
	}

	// possible interesting bios
	// TODO

	// class
	switch pd.class {
	case "Earthlike body", "Water world", "Ammonia world":
		return true
	}

	return false
}
