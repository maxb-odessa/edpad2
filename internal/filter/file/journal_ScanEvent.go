package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/local/sound"
	"edpad2/internal/router"
	"edpad2/pkg/fwt"
	"strings"

	"encoding/json"
	"fmt"
	"time"

	"github.com/danwakefield/fnmatch"
	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"
)

type ScanEvent struct {
	AbsoluteMagnitude     float64 `mapstructure:"AbsoluteMagnitude,omitempty"`
	AgeMy                 int     `mapstructure:"Age_MY,omitempty"`
	Atmosphere            string  `mapstructure:"Atmosphere,omitempty"`
	AtmosphereComposition []struct {
		Name    string  `mapstructure:"Name,omitempty"`
		Percent float64 `mapstructure:"Percent,omitempty"`
	} `mapstructure:"AtmosphereComposition,omitempty"`
	AtmosphereType string  `mapstructure:"AtmosphereType,omitempty"`
	AxialTilt      float64 `mapstructure:"AxialTilt,omitempty"`
	BodyID         int     `mapstructure:"BodyID,omitempty"`
	BodyName       string  `mapstructure:"BodyName,omitempty"`
	Composition    *struct {
		Ice   float64 `mapstructure:"Ice,omitempty"`
		Metal float64 `mapstructure:"Metal,omitempty"`
		Rock  float64 `mapstructure:"Rock,omitempty"`
	} `mapstructure:"Composition,omitempty"`
	DistanceFromArrivalLs float64 `mapstructure:"DistanceFromArrivalLS,omitempty"`
	Eccentricity          float64 `mapstructure:"Eccentricity,omitempty"`
	Landable              bool    `mapstructure:"Landable,omitempty"`
	Luminosity            string  `mapstructure:"Luminosity,omitempty"`
	MassEm                float64 `mapstructure:"MassEM,omitempty"`
	Materials             []struct {
		Name    string  `mapstructure:"Name,omitempty"`
		Percent float64 `mapstructure:"Percent,omitempty"`
	} `mapstructure:"Materials,omitempty"`
	OrbitalInclination float64 `mapstructure:"OrbitalInclination,omitempty"`
	OrbitalPeriod      float64 `mapstructure:"OrbitalPeriod,omitempty"`
	Parents            []struct {
		Null   int `mapstructure:"Null,omitempty"`
		Planet int `mapstructure:"Planet,omitempty"`
		Ring   int `mapstructure:"Ring,omitempty"`
		Star   int `mapstructure:"Star,omitempty"`
	} `mapstructure:"Parents,omitempty"`
	Periapsis    float64 `mapstructure:"Periapsis,omitempty"`
	PlanetClass  string  `mapstructure:"PlanetClass,omitempty"`
	Radius       float64 `mapstructure:"Radius,omitempty"`
	ReserveLevel string  `mapstructure:"ReserveLevel,omitempty"`
	Rings        []struct {
		InnerRad  json.Number `mapstructure:"InnerRad,omitempty"`
		MassMt    json.Number `mapstructure:"MassMT,omitempty"`
		Name      string      `mapstructure:"Name,omitempty"`
		OuterRad  float64     `mapstructure:"OuterRad,omitempty"`
		RingClass string      `mapstructure:"RingClass,omitempty"`
	} `mapstructure:"Rings,omitempty"`
	RotationPeriod     float64   `mapstructure:"RotationPeriod,omitempty"`
	ScanType           string    `mapstructure:"ScanType,omitempty"`
	SemiMajorAxis      float64   `mapstructure:"SemiMajorAxis,omitempty"`
	StarSystem         string    `mapstructure:"StarSystem,omitempty"`
	StarType           string    `mapstructure:"StarType,omitempty"`
	StellarMass        float64   `mapstructure:"StellarMass,omitempty"`
	Subclass           int       `mapstructure:"Subclass,omitempty"`
	SurfaceGravity     float64   `mapstructure:"SurfaceGravity,omitempty"`
	SurfacePressure    float64   `mapstructure:"SurfacePressure,omitempty"`
	SurfaceTemperature float64   `mapstructure:"SurfaceTemperature,omitempty"`
	SystemAddress      int       `mapstructure:"SystemAddress,omitempty"`
	TerraformState     string    `mapstructure:"TerraformState,omitempty"`
	TidalLock          bool      `mapstructure:"TidalLock,omitempty"`
	Volcanism          string    `mapstructure:"Volcanism,omitempty"`
	WasDiscovered      bool      `mapstructure:"WasDiscovered,omitempty"`
	WasMapped          bool      `mapstructure:"WasMapped,omitempty"`
	Event              string    `mapstructure:"event,omitempty"`
	Timestamp          time.Time `mapstructure:"timestamp,omitempty"`
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
func (h *handler) parseStar(ev *ScanEvent) {

	switch ev.StarType[0:1] {
	case "N", "D", "H":
		h.connector.ToRouterCh <- &router.Message{
			Dst: router.LocalSound,
			Data: &sound.Track{
				Id: sound.WARP,
			},
		}
	}

	sd := new(starData)

	sd.id = ev.BodyID

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
	t.Header(&fwt.Header{Text: "Class", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Lum", FgColor: "gray", Underline: true, Italic: true})
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
			t.Cell(idx, &fwt.Cell{Text: "+", FgColor: "white", Bold: true})
		} else {
			t.Cell(idx, &fwt.Cell{Text: ""})
		}

		sname, scolor := CB(s.class)
		t.Cell(idx, &fwt.Cell{Text: fmt.Sprintf("%s %d", sname, s.subClass), FgColor: scolor, Left: true, Bold: true})
		t.Cell(idx, &fwt.Cell{Text: s.luminosity, FgColor: scolor, Left: true})

		t.Cell(idx, &fwt.Cell{Text: formatLargeNum(s.distance)})

		if s.discovered {
			t.Cell(idx, &fwt.Cell{Text: "y", FgColor: "yellow"})
		} else {
			t.Cell(idx, &fwt.Cell{Text: "n", FgColor: "gray"})
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

	text := "\n" + t.Text()
	slog.Debug(99, "STAR:%s", text)
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
		pd.id = ev.BodyID
		pd.parentStarId = getParentStarId(ev)
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

	pd.bios = noteBios(ev)

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

		if !h.remarkablePlanet(p) {
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

		if p.landable {
			t.Cell(idx, &fwt.Cell{Text: "y"})
		} else {
			t.Cell(idx, &fwt.Cell{Text: "-"})
		}

		if p.terraformable {
			t.Cell(idx, &fwt.Cell{Text: "y"})
		} else {
			t.Cell(idx, &fwt.Cell{Text: "-"})
		}

		t.Cell(idx, &fwt.Cell{Text: calcSignals(p.signals), NoFormat: true})

		// will add extra column, could be seen by scrolling window left
		if len(p.bios) > 0 {
			t.Cell(idx, &fwt.Cell{Text: strings.Join(p.bios, ","), NoFormat: true})
		}

		idx++

	}

	if idx > 0 {
		text := "\n" + t.Text()
		slog.Debug(5, "PLANET:%s", text)
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
	}

	return
}

func calcSignals(bs *bodySignals) string {

	sig := ""

	if bs.biological > 0 {
		if bs.biological >= 5 {
			sig += fmt.Sprintf(`<span color="#90ff90">%d</span>`, bs.biological)
		} else {
			sig += fmt.Sprintf(`<span color="#50dd50">%d</span>`, bs.biological)
		}
	} else {
		sig += "-"
	}

	if bs.geological > 0 {
		if bs.geological >= 5 {
			sig += fmt.Sprintf(`<span color="#ffff90">%d</span>`, bs.geological)
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

func (h *handler) remarkablePlanet(pd *planetData) bool {

	// landable + high G
	if pd.landable && pd.gravityG >= float64(sconf.Float32Def("ed journal", "min planet gravity", 1)) {
		return true
	}

	// many rings
	if pd.rings >= int(sconf.Int32Def("ed journal", "min rings num", 3)) {
		return true
	}

	// wide ring
	if pd.ringRad > float64(sconf.Float32Def("ed journal", "min outer ring radius", 25.0)) {
		return true
	}

	if pd.terraformable && sconf.BoolDef("ed journal", "want terraformable", false) {
		return true
	}

	// possible interesting signals
	if pd.signals.biological >= int(sconf.Int32Def("ed journal", "min bio signals", 1)) {
		return true
	}

	if pd.signals.geological >= int(sconf.Int32Def("ed journal", "min geo signals", 1)) {
		return true
	}

	if pd.signals.human >= int(sconf.Int32Def("ed journal", "min human signals", 1)) {
		return true
	}

	if pd.signals.guardian >= int(sconf.Int32Def("ed journal", "min guardian signals", 1)) {
		return true
	}

	if pd.signals.other >= int(sconf.Int32Def("ed journal", "min other signals", 1)) {
		return true
	}

	// class
	wantBodies := sconf.StrDef("ed journal", "want bodies", "Earth*,Water*,Ammonia*,Helium*")
	for _, body := range strings.Split(wantBodies, ",") {
		if fnmatch.Match(strings.TrimSpace(body), pd.class, 0) {
			return true
		}
	}

	return false
}

// TODO: where to use? after planet parse and send to NOTES? as a popup in PLANETS? Where?
func noteBios(ev *ScanEvent) []string {

	var bios []string
	//var starClass string
	var starLum int
	/*
		// find parents star class
		psid := getParentStarId(ev)
		for _, cs := range CurrentSystemStars {
			if cs.id == psid {
				starClass = cs.class
				starLum = luminToInt(cs.luminosity)
				break
			}
		}

		// hmm... not scanned yet?
		if starClass == "" {
			return bios
		}
	*/
	/* TODO map of plant names and their conditions, like
	"aleoida spica" = {
		temp[2] = {120, 160},
		grav[2] = {01, 03},
		atmo[] = {"ammonia", "thin neon", "nitro"}
		geo[] = {"volcanic", "geysers"}
		star[] = {"B, "O, "K", "F"}
		starLum[] = {5, 3, 2}
		body[] = {"hmc", "icy", "rocky"}
	}
	 use fnmatch() for strings matching
	*/

	slog.Debug(999, "%s", starLum)

	return bios
}

func luminToInt(lum string) int {

	if lum[0:4] == "VIII" {
		return 8
	} else if lum[0:3] == "VII" {
		return 7
	} else if lum[0:2] == "VI" {
		return 6
	} else if lum[0:2] == "IX" {
		return 9
	} else if lum[0:2] == "IV" {
		return 4
	} else if lum[0:1] == "V" {
		return 5
	} else if lum[0:3] == "III" {
		return 3
	} else if lum[0:2] == "II" {
		return 2
	} else if lum[0:1] == "I" {
		return 1
	}

	return 0
}

func getParentStarId(ev *ScanEvent) int {

	parent := 0

	for _, p := range ev.Parents {
		if p.Star > parent {
			parent = p.Star
		}
	}

	return parent
}
