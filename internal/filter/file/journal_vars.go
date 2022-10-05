package file

var CurrentSystemName string
var CurrentMainStarName string
var CurrentMainStarClass string
var NextJumpSystem string
var NextJumpStarClass string

type starData struct {
	id           int
	discovered   bool
	isMain       bool
	class        string
	subClass     int
	distanceLs   float64
	luminosity   string
	massSol      float64
	radiusSol    float64
	temperatureK float64
	rings        int
	ringRad      float64
}

type bodySignals struct {
	biological int
	geological int
	human      int
	guardian   int
	other      int
}

type planetData struct {
	id             int
	parentStarId   int
	bodyName       string
	distanceLs     float64
	discovered     bool
	mapped         bool
	class          string
	massEm         float64
	radiusEr       float64
	radiusLs       float64
	gravityG       float64
	temperatureK   float64
	rings          int
	ringRad        float64
	atmosphere     string
	atmosphereType string
	volcanism      string
	landable       bool
	smAxisLs       float64
	rotPeriod      float64
	orbitalPeriod  float64
	terraformable  bool
	signals        *bodySignals
	bios           []string
}

type signalData struct {
	dummy bool
}

var FuelLevel float64

// no just "string" because we must filter out duplicates
var CurrentSystemStars map[string]*starData
var CurrentSystemPlanets map[string]*planetData
var CurrentSystemSignals map[string]*signalData

func resetCurrentSystemStars() {
	CurrentSystemStars = make(map[string]*starData)
}

func resetCurrentSystemPlanets() {
	CurrentSystemPlanets = make(map[string]*planetData)
}

func resetCurrentSystemSignals() {
	CurrentSystemSignals = make(map[string]*signalData)
}

func init() {
	resetCurrentSystemStars()
	resetCurrentSystemPlanets()
	resetCurrentSystemSignals()
}
