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
			t.Cell(idx, &fwt.Cell{Text: "Y", FgColor: "yellow"})
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
	pd.distance = ev.DistanceFromArrivalLs
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

	pd.atmosphere = ev.Atmosphere
	pd.atmosphereType = ev.AtmosphereType
	pd.volcanism = ev.Volcanism

	//pd.bios = possibleBios(pd)

	// show planets list
	h.refreshPlanets()

}

func (h *handler) refreshPlanets() {

	t := fwt.Table{
		Default:   "-",
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

		t.Cell(idx, &fwt.Cell{Text: formatLargeNum(p.distance)})

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

		// those below will add extra columns, could be seen by scrolling window left
		if p.atmosphereType != "" && p.atmosphereType != "none" {
			t.Cell(idx, &fwt.Cell{Text: "Atmo:" + p.atmosphereType + ";", NoFormat: true, Italic: true})
		}

		// this must be recalculated each time
		if p.signals.biological > 0 {
			p.bios = possibleBios(p)
			if len(p.bios) > 0 {
				t.Cell(idx, &fwt.Cell{Text: "Bios:" + strings.Join(p.bios, ",") + ";", NoFormat: true, Italic: true})
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

// TODO: where to use? after planet parse and send to NOTES? as a popup in PLANETS? Where?
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
		slog.Debug(1, "possibleBios(): no parent star found for id pd.parentStarId")
		return []string{}
	}

	return matchBios(pd, sd, planets)
}

func matchBios(pd *planetData, sd *starData, planets []string) []string {

	temp := pd.temperatureK
	grav := pd.gravityG
	atmo := pd.atmosphere
	volc := pd.volcanism
	ptype := pd.class
	dist := pd.distance
	geoSigs := pd.signals.geological

	sclass := sd.class
	slumin := sd.luminosity

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
}

//
var bioDataLimits = map[string]bioLimits{

	"Aleoida": {
		temp:       [2]float64{0.0, 195.0},
		grav:       [2]float64{0.0, 0.27},
		atmos:      []string{"*thin carbon dioxide*", "*thin ammonia*"},
		volcs:      []string{"*"},
		ptypes:     []string{"high metal *", "rocky body"},
		sclass:     []string{"B", "A", "F", "K", "M", "L", "T", "TTS", "Y", "N*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"Amphora": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{""},
		volcs:      []string{"*"},
		ptypes:     []string{"metal rich *"},
		sclass:     []string{"A"},
		slumins:    []string{"*"},
		needBodies: []string{"Earth*", "Ammonia*", "* based life", "water giant"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"Anemone": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{""},
		volcs:      []string{"*"},
		ptypes:     []string{"*"},
		sclass:     []string{"O", "B", "A"},
		slumins:    []string{"*"},
		needBodies: []string{"Earth*", "Ammonia*", "* based life", "water giant"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
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
	},

	"(Brain Tree)": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{""},
		volcs:      []string{"* volcanism"},
		ptypes:     []string{"*"},
		sclass:     []string{"*"},
		slumins:    []string{"*"},
		needBodies: []string{"water giant*", "earth*", "* water based life"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"Cactoida": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.27},
		atmos:      []string{"*ammonia*", "*carbon dioxidie*", "*water*"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body", "high metal *"},
		sclass:     []string{"A", "F", "G", "M", "L", "T", "TTS", "N*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"Clypeus": {
		temp:       [2]float64{0.0, 190.0},
		grav:       [2]float64{0.0, 0.27},
		atmos:      []string{"*thin carbon dioxidie*", "*water*"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body", "high metal *"},
		sclass:     []string{"A", "F", "G", "K", "M", "L", "N*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"Concha": {
		temp:       [2]float64{0.0, 190.0},
		grav:       [2]float64{0.0, 0.27},
		atmos:      []string{"*ammonia*", "*nitrogen*", "*water*"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body", "high metal *"},
		sclass:     []string{"A", "F", "G", "K", "M", "L", "N*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"(Crystalline Shards)": {
		temp:       [2]float64{0.0, 273.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{""},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body", "high metal *"},
		sclass:     []string{"A", "F", "G", "K", "M", "S"},
		slumins:    []string{"*"},
		needBodies: []string{"*ammonia*", "water giant*", "earth*", "* water based life"},
		distLs:     [2]float64{12000.0, 99999999.0},
		needGeo:    false,
	},

	"(Electricae)": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.27},
		atmos:      []string{"*thin helium*", "*argon*", "*neon*"},
		volcs:      []string{"*"},
		ptypes:     []string{"icy body"},
		sclass:     []string{"*"},
		slumins:    []string{"V", "IV", "III", "II", "I"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"Fonticulua": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.27},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"icy *", "rocky ice *"},
		sclass:     []string{"B", "A", "F", "G", "K", "M", "L", "Y", "D*", "N*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"Frutexa": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.27},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body*", "high metal *"},
		sclass:     []string{"B", "F", "G", "M", "L", "TTS", "D*", "N*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"Fumerola": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.27},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*?"},
		ptypes:     []string{"rocky ice *", "rocky body*", "high metal *"},
		sclass:     []string{"*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    true,
	},

	"Fungoida": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.27},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky ice *", "rocky body*", "high metal *"},
		sclass:     []string{"*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"Osseus": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.27},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky ice *", "rocky body*", "high metal *"},
		sclass:     []string{"A", "F", "G", "K", "T", "TTS"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"Recepta": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.27},
		atmos:      []string{"*thin sulphur dioxide*"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky ice *", "rocky body*", "high metal *"},
		sclass:     []string{"A", "F", "G", "K", "M", "T"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
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
	},

	"Stratum": {
		temp:       [2]float64{0.0, 190.0},
		grav:       [2]float64{0.0, 9999.0},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body*", "high metal *"},
		sclass:     []string{"*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"Tubus": {
		temp:       [2]float64{0.0, 190.0},
		grav:       [2]float64{0.0, 0.15},
		atmos:      []string{"*thin *carbon*", "*thin *ammonia*"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body*", "high metal *"},
		sclass:     []string{"B", "A", "F", "G", "K", "M", "L", "T", "TTS", "NS"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},

	"Tussock": {
		temp:       [2]float64{0.0, 9999.0},
		grav:       [2]float64{0.0, 0.27},
		atmos:      []string{"*thin *"},
		volcs:      []string{"*"},
		ptypes:     []string{"rocky body*"},
		sclass:     []string{"F", "G", "K", "M", "L", "T", "D*"},
		slumins:    []string{"*"},
		needBodies: []string{"*"},
		distLs:     [2]float64{0.0, 99999999.0},
		needGeo:    false,
	},
}
