package file

var CurrentSystemName string
var CurrentMainStarName string
var CurrentMainStarClass string

type starData struct {
	discovered   bool
	isMain       bool
	class        string
	subClass     int
	luminosity   string
	massSol      float64
	radiusSol    float64
	temperatureK string
	hasBelt      bool
}

type planetData struct {
	discovered  bool
	scanned     bool
	class       string
	massEm      float64
	radiusEm    float64
	gravity     float64
	rings       int
	wideRing    bool
	possibleBio []string
}

type signalData struct {
	dummyfornow string
}

var CurrentSystemStars map[string]starData
var CurrentSystemPlanets map[string]planetData
var CurrentSystemSignals map[string]signalData

func resetCurrentSystemStars() {
	CurrentSystemStars = make(map[string]starData)
}

func resetCurrentSystemPlanets() {
	CurrentSystemPlanets = make(map[string]planetData)
}

func resetCurrentSystemSignals() {
	CurrentSystemSignals = make(map[string]signalData)
}

func init() {
	resetCurrentSystemStars()
	resetCurrentSystemPlanets()
	resetCurrentSystemSignals()
}
