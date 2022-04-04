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

	if ev.StarType != "" {
		h.parseStar(ev)
	} else if ev.PlanetClass != "" {
		h.parsePlanet(ev)
	} else {
		// ateroids? other?
	}
}

func formatE(val float64) string {
	if val >= 100.0 { // yes, 100.0 is correct
		return fmt.Sprintf("%4.1fK", val/1000.0)
	} else {
		return fmt.Sprintf("%4.2f", val)
	}
}

func (h *handler) parseStar(ev *ScanEvent) {

	sd := new(starData)

	sd.discovered = ev.WasDiscovered
	if ev.BodyName == CurrentMainStarName {
		sd.isMain = true
	}
	sd.class = CB(ev.StarType, 4)
	sd.subClass = ev.Subclass
	sd.distance = ev.DistanceFromArrivalLs
	sd.luminosity = ev.Luminosity
	sd.rings = len(ev.Rings)
	sd.massSol = ev.StellarMass
	sd.radiusSol = ev.Radius / SOLAR_RADIUS
	sd.temperatureK = formatLargeNum(ev.SurfaceTemperature)
	sd.rings, sd.ringRad = calcRings(ev)

	CurrentSystemStars[ev.BodyName] = sd

	text := "\n" +
		` <i><u><span color="gray">` +
		`     Class   Dst(ls) Disc  Rn   Rr  M(sol) R(sol) Temp(K)` +
		`</span></u></i>` +
		"\n"

	yes := `<span color="yellow">yes</span>`
	no := `<span color="gray">no </span>`

	idx := 0

	for _, s := range CurrentSystemStars {

		mainstar := " "
		if s.isMain {
			mainstar = "*"
		}

		discovered := no
		if s.discovered {
			discovered = yes
		}

		ringRad := "   -"
		ringsNum := " -"
		if s.rings > 0 {
			ringRad = fmt.Sprintf("%4.0f", s.ringRad/LIGHT_SECOND)
			if s.ringRad >= MIN_RING_OUT_RAD {
				ringRad = `<span color="white">` + ringRad + `</span>`
			}
			ringsNum = fmt.Sprintf("%2d", s.rings)
		}

		text += fmt.Sprintf(" %s %-s%-1d %-3.3s  %6.6s  %s   %2.2s %4.4s  %3.2f   %3.2f   %s",
			mainstar,
			s.class,
			s.subClass,
			s.luminosity,
			formatLargeNum(s.distance),
			discovered,
			ringsNum,
			ringRad,
			s.massSol,
			s.radiusSol,
			s.temperatureK)

		idx++

		text += "\n"
	}

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_SYSTEM,
			Text:           text,
			AppendText:     false,
			UpdateText:     true,
			Subtitle:       fmt.Sprintf("[%d]", idx),
			UpdateSubtitle: false,
		},
	}

	return
}

func (h *handler) parsePlanet(ev *ScanEvent) {

	// either add new or update existing one
	pd, ok := CurrentSystemPlanets[ev.BodyName]
	if !ok {
		pd = new(planetData)
		pd.signals = new(bodySignals)
		CurrentSystemPlanets[ev.BodyName] = pd
	}

	//  calc short name from BodyName, like "Graea Hypue FV-X d1-30 2 a" > "2 a"
	bname := "        " + ev.BodyName // a hack to avoid too short names
	pd.shortName = bname[len(bname)-8:]
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
	pd.rings, pd.ringRad = calcRings(ev)

	// show planets list
	h.refreshPlanets()

}

func (h *handler) refreshPlanets() {

	// table headers
	text := "\n" +
		` <i><u><span color="gray">` +
		`  Name    Type  D M  M(e)  R(e)  Grav  T(K)  Rn  Rr  Ld TF bgHGO` +
		`</span></u></i>` +
		"\n"

	idx := 0

	dyes := `<span color="yellow">y</span>`
	yes := `<span color="white">y</span>`
	no := `<span color="gray">n</span>`

	for _, p := range CurrentSystemPlanets {

		// not enuff data yet (i.e. signals only detected, no Scan event happened), skip it
		if p.shortName == "" {
			continue
		}

		if !remarkablePlanet(p) {
			continue
		}

		ringRad := " - "
		ringsNum := " -"
		if p.rings > 0 {
			ringRad = fmt.Sprintf("%3.0f", p.ringRad/LIGHT_SECOND)
			if p.ringRad >= MIN_RING_OUT_RAD {
				ringRad = `<span color="white">` + ringRad + `</span>`
			}
			ringsNum = fmt.Sprintf("%2d", p.rings)
		}

		discovered := no
		if p.discovered {
			discovered = dyes
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

		text += fmt.Sprintf(" ~%-8.8s %s %s %s %s %s %s  %s  %s %s  %s  %s  %s",
			p.shortName,
			CB(p.class, -5),
			discovered,
			mapped,
			formatE(p.massEm),
			formatE(p.radiusEm),
			formatE(p.gravityG),
			formatLargeNum(p.temperatureK),
			ringsNum,
			ringRad,
			landable,
			terraformable,
			calcSignals(p.signals),
		)

		text += "\n"

		idx++

	}

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_PLANETS,
			Text:           text,
			AppendText:     false,
			UpdateText:     true,
			Subtitle:       fmt.Sprintf("[%d]", idx),
			UpdateSubtitle: true,
		},
	}

	return
}

func calcSignals(bs *bodySignals) string {

	sig := ""

	if bs.biological > 0 {
		if bs.biological > 6 {
			sig += fmt.Sprintf(`<span color="#80ff80">%d</span>`, bs.biological)
		} else {
			sig += fmt.Sprintf(`<span color="#50dd50">%d</span>`, bs.biological)
		}
	} else {
		sig += "-"
	}

	if bs.geological > 0 {
		if bs.geological > 6 {
			sig += fmt.Sprintf(`<span color="#ffff80">%d</span>`, bs.geological)
		} else {
			sig += fmt.Sprintf(`<span color="#dddd50">%d</span>`, bs.geological)
		}
	} else {
		sig += "-"
	}

	if bs.human > 0 {
		sig += fmt.Sprintf(`<span color="#3090ff">%d</span>`, bs.human)
	} else {
		sig += "-"
	}

	if bs.guardian > 0 {
		sig += fmt.Sprintf(`<span color="#9030ff">%d"</span>`, bs.guardian)
	} else {
		sig += "-"
	}

	if bs.other > 0 {
		sig += fmt.Sprintf(`<span color="ff3030">%d</span>`, bs.other)
	} else {
		sig += "-"
	}

	return sig
}

func remarkablePlanet(pd *planetData) bool {

	// landable + high G
	if pd.landable && pd.gravityG >= 2.5 {
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

	// possible interesting signals
	if pd.signals.biological+pd.signals.human+pd.signals.guardian+pd.signals.other > 0 {
		return true
	}

	// i don't need geo signals atm
	/*
		if pd.signals.geological > 0 {
			return true
		}
	*/

	// class
	switch pd.class {
	case "Earthlike body", "Water world", "Ammonia world":
		return true
	}

	return false
}

// calc if has wide ring
func calcRings(ev *ScanEvent) (rNum int, wRad float64) {

	for _, r := range ev.Rings {
		if r.Name[len(r.Name)-5:] != " Ring" {
			continue
		}
		rNum++
		if wRad < r.OuterRad {
			wRad = r.OuterRad
		}
	}

	return
}
