package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/local/sound"
	"edpad2/internal/router"
	"edpad2/pkg/fwt"

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

func formatE(val float64) string {
	if val >= 100.0 { // yes, 100.0 is correct
		return fmt.Sprintf("%3.1fK", val/1000.0)
	} else {
		return fmt.Sprintf("%4.1f", val)
	}
}

func (h *handler) parseStar(ev *ScanEvent) {

	switch ev.StarType[0:1] {
	case "N", "D", "H":
		h.connector.ToRouterCh <- &router.Message{
			Dst: router.LocalSound,
			Data: &sound.Track{
				Id: sound.ALARM,
			},
		}
	}

	sd := new(starData)

	sd.discovered = ev.WasDiscovered
	if ev.BodyName == CurrentMainStarName {
		sd.isMain = true
	}
	sd.class = ev.StarType
	sd.subClass = ev.Subclass
	sd.distance = ev.DistanceFromArrivalLs
	sd.luminosity = ev.Luminosity
	sd.rings = len(ev.Rings)
	sd.massSol = ev.StellarMass
	sd.radiusSol = ev.Radius / SOLAR_RADIUS
	sd.temperatureK = ev.SurfaceTemperature
	sd.rings, sd.ringRad = calcRings(ev)

	CurrentSystemStars[ev.BodyName] = sd

	t := &fwt.Table{
		Delimiter: " ",
		Pango:     true,
		Default:   "-",
	}

	t.Header(&fwt.Header{Text: " ", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "  Class   ", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Dist(ls)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Disc", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: " Rn", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: " Rr", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "M(sol)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "R(sol)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Temp(K)", FgColor: "gray", Underline: true, Italic: true})

	idx := 0

	for _, s := range CurrentSystemStars {

		if s.isMain {
			t.Cell(idx, &fwt.Cell{Text: display.STAR4FILLED, FgColor: "white", Bold: true})
		} else {
			t.Cell(idx, &fwt.Cell{Text: ""})
		}
		sname, scolor := CB(s.class)
		t.Cell(idx, &fwt.Cell{Text: fmt.Sprintf("%s %d %s", sname, s.subClass, s.luminosity), FgColor: scolor, Left: true, Bold: true})

		t.Cell(idx, &fwt.Cell{Text: formatLargeNum(s.distance)})

		if s.discovered {
			t.Cell(idx, &fwt.Cell{Text: display.CIRCLEDY, FgColor: "yellow"})
		} else {
			t.Cell(idx, &fwt.Cell{Text: display.CIRCLEDN, FgColor: "gray"})
		}

		ringRad := "-"
		ringsNum := "-"
		if s.rings > 0 {
			ringRad = fmt.Sprintf("%.0f", s.ringRad/LIGHT_SECOND)
			ringsNum = fmt.Sprintf("%d", s.rings)
		}
		t.Cell(idx, &fwt.Cell{Text: ringsNum})
		t.Cell(idx, &fwt.Cell{Text: ringRad})

		t.Cell(idx, &fwt.Cell{Text: fmt.Sprintf("%.2f", s.massSol)})
		t.Cell(idx, &fwt.Cell{Text: fmt.Sprintf("%.2f", s.radiusSol)})
		t.Cell(idx, &fwt.Cell{Text: formatLargeNum(s.temperatureK)})

		idx++

	}

	text := t.Text()
	slog.Debug(0, "STAR:\n%s", text)
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

	pd.bodyName = ev.BodyName
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

	t := fwt.Table{
		Default:   "-",
		Delimiter: " ",
		Pango:     true,
	}

	t.Header(&fwt.Header{Text: "  Name  ", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Type ", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "D", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "M", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "  M(e)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "  R(e)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "  Grav", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: " T(K)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Rn", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: " Rr", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Ld", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Tf", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "bgHGO", FgColor: "gray", Underline: true, Italic: true})

	idx := 0

	for _, p := range CurrentSystemPlanets {

		// not enuff data yet (i.e. signals only detected, no Scan event happened), skip it
		if p.bodyName == "" {
			continue
		}

		if !remarkablePlanet(p) {
			continue
		}

		t.Cell(idx, &fwt.Cell{Text: p.bodyName, Left: true})

		ptype, pcolor := CB(p.class)
		t.Cell(idx, &fwt.Cell{Text: ptype, FgColor: pcolor, Left: true})

		if p.discovered {
			t.Cell(idx, &fwt.Cell{Text: "y", Bold: true, FgColor: "yellow"})
		} else {
			t.Cell(idx, &fwt.Cell{Text: "-", FgColor: "gray"})
		}

		if p.mapped {
			t.Cell(idx, &fwt.Cell{Text: "y", Bold: true, FgColor: "yellow"})
		} else {
			t.Cell(idx, &fwt.Cell{Text: "-", FgColor: "gray"})
		}

		t.Cell(idx, &fwt.Cell{Text: formatE(p.massEm)})
		t.Cell(idx, &fwt.Cell{Text: formatE(p.radiusEm)})
		t.Cell(idx, &fwt.Cell{Text: formatE(p.gravityG)})
		t.Cell(idx, &fwt.Cell{Text: formatLargeNum(p.temperatureK)})

		ringRad := "-"
		ringsNum := "-"
		if p.rings > 0 {
			ringRad = fmt.Sprintf("%3.0f", p.ringRad/LIGHT_SECOND)
			ringsNum = fmt.Sprintf("%d", p.rings)
		}
		t.Cell(idx, &fwt.Cell{Text: ringsNum})
		t.Cell(idx, &fwt.Cell{Text: ringRad})

		if p.terraformable {
			t.Cell(idx, &fwt.Cell{Text: "y"})
		} else {
			t.Cell(idx, &fwt.Cell{Text: "-"})
		}

		if p.landable {
			t.Cell(idx, &fwt.Cell{Text: "y"})
		} else {
			t.Cell(idx, &fwt.Cell{Text: "-"})
		}

		t.Cell(idx, &fwt.Cell{Text: calcSignals(p.signals), NoFormat: true})

		idx++

	}
	tt := t.Text()
	slog.Debug(0, "PLANET: %s", tt)
	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:       display.VP_PLANETS,
			Text:           tt,
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
