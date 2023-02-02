package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"edpad2/pkg/fwt"
	"sort"
	"strings"

	"encoding/json"
	"fmt"
	"math"
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

	sd := new(starData)

	sd.id = ev.BodyID

	sd.discovered = ev.WasDiscovered
	if ev.BodyName == CurrentMainStarName {
		sd.isMain = true
	}
	sd.class = ev.StarType
	sd.subClass = ev.Subclass
	sd.distanceLs = ev.DistanceFromArrivalLs
	sd.luminosity = ev.Luminosity
	sd.rings = len(ev.Rings)
	sd.massSol = ev.StellarMass
	sd.radiusSol = ev.Radius / SOLAR_RADIUS
	sd.temperatureK = ev.SurfaceTemperature
	sd.rings, sd.ringRadLs = calcRings(ev)

	CurrentSystemStars[ev.BodyName] = sd

	t := &fwt.Table{
		Delimiter: " ",
		Pango:     true,
		Default:   "",
	}

	t.Header(&fwt.Header{Text: " ", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Class ", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Lumin ", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Dist(ls)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Disc", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Rn", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Rr(ls)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "M(sol)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "R(sol)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Temp(K)", FgColor: "gray", Underline: true, Italic: true})

	var keys []string
	for key, _ := range CurrentSystemStars {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	idx := 0

	for _, key := range keys {

		s, _ := CurrentSystemStars[key]

		if s.isMain {
			t.Cell(idx, &fwt.Cell{Text: "*", FgColor: "white", Bold: true})
		} else {
			t.Cell(idx, &fwt.Cell{Text: ""})
		}

		sname, scolor := CB(s.class)
		t.Cell(idx, &fwt.Cell{Text: fmt.Sprintf("%s %d", sname, s.subClass), FgColor: scolor, Left: true, Bold: true})
		t.Cell(idx, &fwt.Cell{Text: s.luminosity, FgColor: scolor, Left: true})

		t.Cell(idx, &fwt.Cell{Text: formatLargeNum(s.distanceLs)})

		if s.discovered {
			t.Cell(idx, &fwt.Cell{Text: "Y", FgColor: "yellow"})
		} else {
			t.Cell(idx, &fwt.Cell{Text: "n", FgColor: "gray"})
		}

		ringRadLs := ""
		ringsNum := ""
		if s.rings > 0 {
			ringRadLs = fmt.Sprintf("%.0f", s.ringRadLs/LIGHT_SECOND)
			ringsNum = fmt.Sprintf("%d", s.rings)
		}
		t.Cell(idx, &fwt.Cell{Text: ringsNum})
		t.Cell(idx, &fwt.Cell{Text: ringRadLs})

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
		pd.parentPlanetId = getParentPlanetId(ev)
	}

	pd.bodyName = ev.BodyName
	pd.class = ev.PlanetClass
	pd.distanceLs = ev.DistanceFromArrivalLs
	pd.discovered = ev.WasDiscovered
	pd.mapped = ev.WasMapped
	pd.massEm = ev.MassEm
	pd.rotPeriod = ev.RotationPeriod
	pd.orbitalPeriod = ev.OrbitalPeriod
	pd.radiusEr = ev.Radius / EARTH_RADIUS
	pd.radiusLs = ev.Radius / LIGHT_SECOND
	pd.gravityG = ev.SurfaceGravity / 10.0
	pd.temperatureK = ev.SurfaceTemperature
	pd.rings = len(ev.Rings)
	pd.landable = ev.Landable
	pd.smAxisLs = ev.SemiMajorAxis / LIGHT_SECOND
	if ev.TerraformState != "" {
		pd.terraformable = true
	}
	pd.rings, pd.ringRadLs = calcRings(ev)

	pd.atmosphere = ev.Atmosphere
	pd.atmosphereType = ev.AtmosphereType
	pd.volcanism = ev.Volcanism

	codexText := ""

	// close bodies, close rings, etc
	for pname, pdata := range CurrentSystemPlanets {

		// skip self comparison
		if pd.bodyName == pname {
			continue
		}

		centerDistLs := math.Abs(pd.distanceLs - pdata.distanceLs)
		surfaceDistLs := math.Abs(centerDistLs - pd.radiusLs - pdata.radiusLs)
		ringDistLs := math.Abs(centerDistLs - pd.ringRadLs - pdata.ringRadLs)
		maxRad := math.Max(pd.radiusLs, pdata.radiusLs)
		wantedRadDistLs := maxRad * float64(sconf.Float32Def("ed journal", "max bodies distance", 3.0))
		smAxDistLs := pd.smAxisLs + pdata.smAxisLs
		smAxSurfDistLs := smAxDistLs - pd.radiusLs - pdata.radiusLs
		smAxRingDistLs := smAxDistLs - pd.ringRadLs - pdata.ringRadLs

		if surfaceDistLs <= wantedRadDistLs && smAxSurfDistLs <= wantedRadDistLs {
			codexText += fmt.Sprintf("Close BODIES, approx distance: %.4f Ls\n"+
				` |- SemiAxis: %.4f Ls`+"\n"+
				` |- R: %.4f Ls, %s`+"\n"+
				` |- R: %.4f Ls, %s`+"\n",
				smAxSurfDistLs,
				smAxDistLs,
				pd.radiusLs,
				pd.bodyName,
				pdata.radiusLs,
				pdata.bodyName,
			)
		}

		if ringDistLs <= wantedRadDistLs && smAxRingDistLs <= wantedRadDistLs {
			codexText += fmt.Sprintf("Close RING(S), approx distance: %.4f Ls\n"+
				` |- SemiAxis: %.4f Ls`+"\n"+
				` |- R: %.4f Ls, %s`+"\n"+
				` |- R: %.4f Ls, %s`+"\n",
				smAxRingDistLs,
				smAxDistLs,
				pd.radiusLs,
				pd.bodyName,
				pdata.radiusLs,
				pdata.bodyName,
			)
		}

		// prev scanned body has parent body and it is US
		var parent, satellite *planetData
		if pdata.parentPlanetId > 0 && pdata.parentPlanetId == pd.id {
			parent = pdata
			satellite = pd
			// current body has parent and it was already scanned
		} else if pd.parentPlanetId > 0 && pd.parentPlanetId == pdata.id {
			parent = pd
			satellite = pdata
		}

		// body orbit is inside parent's ring
		if parent != nil && satellite != nil && parent.ringRadLs/LIGHT_SECOND > satellite.smAxisLs {
			codexText += fmt.Sprintf("Body orbit is inside parent ring\n"+
				` |- Ring radius: %.4f Ls, Orbit radius: %.4f Ls`+"\n"+
				` |- Parent:    %s`+"\n"+
				` |- Sattelite: %s`+"\n",
				parent.ringRadLs/LIGHT_SECOND,
				satellite.smAxisLs,
				parent.bodyName,
				satellite.bodyName,
			)
		}

	}

	// short orbital period
	orbPeriod := math.Abs(pd.orbitalPeriod / SECONDS_IN_HOUR)
	if orbPeriod <= float64(sconf.Float32Def("ed journal", "max orbital period", 3.0)) {
		codexText += fmt.Sprintf("Short orbital period = %.2f Hours\n"+
			" |- %s\n",
			orbPeriod, pd.bodyName)
	}

	// fast rotating
	rotPeriod := math.Abs(pd.rotPeriod) / SECONDS_IN_HOUR
	if ev.TidalLock == false && rotPeriod <= float64(sconf.Float32Def("ed journal", "max rotation period", 3.0)) {
		codexText += fmt.Sprintf("Fast rotation period = %.2f Hours\n"+
			" |- %s\n",
			rotPeriod, pd.bodyName)

	}

	if codexText != "" {
		h.connector.ToRouterCh <- &router.Message{
			Dst: router.LocalDisplay,
			Data: &display.Text{
				ViewPort:       display.VP_LOGS,
				Text:           codexText,
				AppendText:     true,
				UpdateText:     true,
				Subtitle:       "[!]",
				UpdateSubtitle: true,
			},
		}
	}

	// show planets list
	h.refreshPlanets()

}

func (h *handler) refreshPlanets() {

	t := fwt.Table{
		Default:   "",
		Delimiter: " ",
		Pango:     true,
	}

	t.Header(&fwt.Header{Text: " Name ", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Type ", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "  Dist", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "D/M", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: " M(e)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: " R(e)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: " Grav", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: " T(K)", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Rn", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: " Rr", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "L/T", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "bgHGO", FgColor: "gray", Underline: true, Italic: true})
	t.Header(&fwt.Header{Text: "Atmo  ", FgColor: "gray", Underline: true, Italic: true})

	var keys []string
	for key, _ := range CurrentSystemPlanets {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	idx := 0
	for _, key := range keys {

		p, _ := CurrentSystemPlanets[key]

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

		t.Cell(idx, &fwt.Cell{Text: formatLargeNum(p.distanceLs)})

		discovered := "n"
		mapped := "n"
		if p.discovered {
			discovered = "Y"
		}

		if p.mapped {
			mapped = "Y"
		}

		if p.discovered || p.mapped {
			t.Cell(idx, &fwt.Cell{Text: discovered + "/" + mapped, Bold: true, FgColor: "yellow"})
		} else {
			t.Cell(idx, &fwt.Cell{Text: "n/n", FgColor: "gray"})
		}

		t.Cell(idx, &fwt.Cell{Text: formatE(p.massEm)})
		t.Cell(idx, &fwt.Cell{Text: formatE(p.radiusEr)})
		t.Cell(idx, &fwt.Cell{Text: formatE(p.gravityG)})
		t.Cell(idx, &fwt.Cell{Text: formatLargeNum(p.temperatureK)})

		ringRadLs := ""
		ringsNum := ""
		if p.rings > 0 {
			ringRadLs = fmt.Sprintf("%3.0f", p.ringRadLs/LIGHT_SECOND)
			ringsNum = fmt.Sprintf("%d", p.rings)
		}
		t.Cell(idx, &fwt.Cell{Text: ringsNum})
		t.Cell(idx, &fwt.Cell{Text: ringRadLs})

		landable := "n"
		terraformable := "n"
		if p.landable {
			landable = "Y"
		}

		if p.terraformable {
			terraformable = "Y"
		}

		if p.landable || p.terraformable {
			t.Cell(idx, &fwt.Cell{Text: landable + "/" + terraformable, FgColor: "white"})
		} else {
			t.Cell(idx, &fwt.Cell{Text: landable + "/" + terraformable, FgColor: "gray"})
		}

		t.Cell(idx, &fwt.Cell{Text: calcSignals(p.signals), NoFormat: true})

		t.Cell(idx, &fwt.Cell{Text: atmoFormula(p.atmosphereType), FgColor: "#5050AA", Left: true, Italic: true})

		// this must be recalculated each time
		if p.signals.biological > 0 {
			p.bios = possibleBios(p)
			if len(p.bios) > 0 {
				sort.Strings(p.bios)
				// this will add extra columns, could be seen by scrolling window left
				t.Cell(idx, &fwt.Cell{Text: strings.Join(p.bios, ","), FgColor: "#50AA50", NoFormat: true, Italic: true})
			}
		}

		idx++

	}

	if idx > 0 {
		text := "\n" + t.Text()
		slog.Debug(99, "PLANET:%s", text)
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
	if pd.ringRadLs/LIGHT_SECOND > float64(sconf.Float32Def("ed journal", "min outer ring radius", 15.0)) {
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
		if fnmatch.Match(strings.TrimSpace(body), pd.class, fnmatch.FNM_IGNORECASE) {
			return true
		}
	}

	// atmos
	if pd.landable {
		wantAtmos := sconf.StrDef("ed journal", "want atmospheres", "none")
		for _, atmo := range strings.Split(wantAtmos, ",") {
			if fnmatch.Match(strings.TrimSpace(atmo), pd.atmosphere, fnmatch.FNM_IGNORECASE) {
				return true
			}
		}
	}

	return false
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

func getParentPlanetId(ev *ScanEvent) int {

	parent := 0

	for _, p := range ev.Parents {
		if p.Planet > parent {
			parent = p.Planet
		}
	}

	return parent
}

// TODO: where to use? after planet parse and send to LOGS? as a popup in PLANETS? Where?
func possibleBios(pd *planetData) []string {

	var sd *starData
	var planets []string

	// find parent star
	for _, s := range CurrentSystemStars {
		if s.id == pd.parentStarId {
			sd = s
			break
		}
	}

	// baricenter! find lowest star id to have at least something
	if sd == nil {
		lowestId := 99999
		for _, s := range CurrentSystemStars {
			if s.id < lowestId {
				lowestId = s.id
				sd = s
			}
		}
	}

	// get all discovered planets list
	for _, p := range CurrentSystemPlanets {
		planets = append(planets, p.class)
	}

	if sd == nil {
		slog.Debug(1, "possibleBios(): no parent star found for id = %d", pd.parentStarId)
		//return []string{}
	}

	return matchBios(pd, sd, planets)
}

func matchBios(pd *planetData, sd *starData, planets []string) []string {

	temp := pd.temperatureK
	grav := pd.gravityG
	atmo := pd.atmosphere
	volc := pd.volcanism
	ptype := pd.class
	dist := pd.distanceLs
	geoSigs := pd.signals.geological

	var sclass, slumin string
	if sd != nil {
		sclass = sd.class
		slumin = sd.luminosity
	} else {
		sclass = "none"
		slumin = "none"
	}

	var bios []string

	for name, bl := range bioDataLimits {

		if temp < bl.temp[0] || temp > bl.temp[1] {
			continue
		}

		if grav < bl.grav[0] || grav > bl.grav[1] {
			continue
		}

		if !matches(atmo, bl.atmos) {
			continue
		}

		if !matches(volc, bl.volcs) {
			continue
		}

		if !matches(ptype, bl.ptypes) {
			continue
		}

		if !matches(sclass, bl.sclass) {
			continue
		}

		if !matches(slumin, bl.slumins) {
			continue
		}

		if !matchesAny(planets, bl.needBodies) {
			continue
		}

		if dist < bl.distLs[0] || dist > bl.distLs[1] {
			continue
		}

		if bl.needGeo && geoSigs == 0 {
			continue
		}

		bios = append(bios, name)

	}

	return bios
}

func matches(what string, where []string) bool {

	for _, wp := range where {
		if fnmatch.Match(wp, what, fnmatch.FNM_IGNORECASE) {
			return true
		}
	}

	return false
}

func matchesAny(what []string, where []string) bool {

	for _, wha := range what {
		if matches(wha, where) {
			return true
		}
	}

	return false
}

type bioLimits struct {
	temp       [2]float64
	grav       [2]float64
	atmos      []string
	volcs      []string
	ptypes     []string
	sclass     []string
	slumins    []string
	needBodies []string
	distLs     [2]float64
	needGeo    bool
	diversityM int
}

var bioDataLimits = map[string]bioLimits{

	"Aleoida": {
		temp:       [2]float64{0.0, 195.0},
		grav:       [2]float64{0.0, 0.28},
		atmos:      []string{"*thin carbon dioxide*", "*thin ammonia*"},
		volcs:      []string{"*"},
		ptypes:     []string{"*high metal*", "rocky body"},
		sclass:     []string{"B", "A", "F", "K", "M", "L", "T*", "Y", "N*", "H"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 150,
	},

	"Amphora": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{""},
		volcs:      []string{"*"},
		ptypes:     []string{"metal rich *"},
		sclass:     []string{"A"},
		slumins:    []string{"*"},
		needBodies: []string{"Earth*", "Ammonia*", "*based life", "water giant"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 100,
	},

	"Anemone": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{"", "*thin*"},
		volcs:      []string{"*"},
		ptypes:     []string{"*"},
		sclass:     []string{"O", "B", "A", "H"},
		slumins:    []string{"*"},
		needBodies: []string{"Earth*", "Ammonia*", "*based life", "water giant"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 100,
	},

	"(Bark Mounds)": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{""},
		volcs:      []string{"*"},
		ptypes:     []string{"*"},
		sclass:     []string{"*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 100,
	},

	"Bacterium": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"*"},
		sclass:     []string{"*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 500,
	},

	"(Brain Tree)": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{"*"},
		volcs:      []string{"* volcanism*"},
		ptypes:     []string{"rocky ice*", "rocky body", "metal rich*", "high metal *"},
		sclass:     []string{"*"},
		slumins:    []string{"*"},
		needBodies: []string{"*water giant*", "*earth*", "* water based life*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 100,
	},

	"Cactoida": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.28},
		atmos:      []string{"*ammonia*", "*carbon dioxide*", "*water*"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body", "high metal *"},
		sclass:     []string{"A", "F", "G", "M", "L", "T*", "N*", "H"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 300,
	},

	"Clypeus": {
		temp:       [2]float64{190.0, 9999.0},
		grav:       [2]float64{0.0, 0.28},
		atmos:      []string{"*thin carbon dioxide*", "*water*"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body", "high metal *"},
		sclass:     []string{"A", "F", "G", "K", "M", "L", "N*", "H"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 150,
	},

	"Concha": {
		temp:       [2]float64{0.0, 190.0},
		grav:       [2]float64{0.0, 0.28},
		atmos:      []string{"*carbon dioxide*", "*ammonia*", "*nitrogen*", "*water*"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body", "high metal *"},
		sclass:     []string{"A", "F", "G", "K", "M", "L", "N*", "H"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 150,
	},

	"(Crystalline Shards)": {
		temp:       [2]float64{0.0, 273.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{""},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body", "high metal *"},
		sclass:     []string{"A", "F", "G", "K", "M", "S", "H"},
		slumins:    []string{"*"},
		needBodies: []string{"*ammonia*", "water giant*", "earth*", "* water based life"},
		distLs:     [2]float64{12000.0, 99999999.0},
		needGeo:    false,
		diversityM: 100,
	},

	"(Electricae)": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.28},
		atmos:      []string{"*helium*", "*argon*", "*neon*"},
		volcs:      []string{"*"},
		ptypes:     []string{"icy body"},
		sclass:     []string{"*"},
		slumins:    []string{"VII", "VI", "V", "IV", "III", "II", "I"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 1000,
	},

	"Fonticulua": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.28},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"icy *", "rocky ice *"},
		sclass:     []string{"B", "A", "F", "G", "K", "M", "L", "Y", "D*", "N*", "H", "T*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 500,
	},

	"Frutexa": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.28},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body*", "high metal *"},
		sclass:     []string{"B", "F", "G", "M", "L", "T*", "D*", "N*", "H"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 150,
	},

	"Fumerola": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.28},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*?"},
		ptypes:     []string{"icy body", "rocky ice *", "rocky body*", "high metal *"},
		sclass:     []string{"*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    true,
		diversityM: 100,
	},

	"Fungoida": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.28},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky ice *", "rocky body*", "high metal *"},
		sclass:     []string{"*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 300,
	},

	"Osseus": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.28},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky ice *", "rocky body*", "high metal *"},
		sclass:     []string{"A", "F", "G", "K", "T*", "H"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 800,
	},

	"Recepta": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.28},
		atmos:      []string{"*thin sulphur dioxide*"},
		volcs:      []string{"*"},
		ptypes:     []string{"*"},
		sclass:     []string{"A", "F", "G", "K", "M", "T", "H"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 150,
	},

	"Sinuous Tubers": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{"*"},
		volcs:      []string{"* volcan*"},
		ptypes:     []string{"metal rich*", "rocky ice *", "rocky body*", "high metal *"},
		sclass:     []string{"*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 100,
	},

	"Stratum": {
		temp:       [2]float64{165.0, 9999.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body*", "high metal *"},
		sclass:     []string{"*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 500,
	},

	"Tubus": {
		temp:       [2]float64{160.0, 9999.0},
		grav:       [2]float64{0.0, 0.16},
		atmos:      []string{"*thin *carbon*", "*thin *ammonia*"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body*", "high metal *"},
		sclass:     []string{"B", "A", "F", "G", "K", "M", "L", "T*", "N*", "H"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 800,
	},

	"Tussock": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.28},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body*", "rocky ice*"},
		sclass:     []string{"F", "G", "K", "M", "L", "T", "D*", "H"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
		diversityM: 200,
	},
}

func atmoFormula(atmo string) string {

	switch atmo {
	case "Ammonia", "AmmoniaRich":
		return "NH3"
	case "AmmoniaOxygen":
		return "NH3+02"
	case "Argon", "ArgonRich":
		return "Ar"
	case "CarbonDioxide", "CarbonDioxideRich":
		return "CO2"
	case "EarthLike":
		return "N2+O2"
	case "Helium":
		return "He"
	case "MetallicVapour":
		return "Me+"
	case "Methane", "MethaneRich":
		return "CH4"
	case "Neon", "NeonRich":
		return "Ne"
	case "Nitrogen":
		return "N2"
	case "None":
		return ""
	case "Oxygen":
		return "O2"
	case "SilicateVapour":
		return "SiO4+"
	case "SulphurDioxide":
		return "SO2"
	case "Water", "WaterRich":
		return "H2O"
	}

	return atmo
}
