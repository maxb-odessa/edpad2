package file

var CurrentSystemName string
var CurrentMainStarName string
var CurrentMainStarClass string

type starData struct {
	discovered   bool
	isMain       bool
	class        string
	subClass     int
	distance     float64
	luminosity   string
	massSol      float64
	radiusSol    float64
	temperatureK string
	hasBelt      bool
}

type planetData struct {
	discovered    bool
	mapped        bool
	class         string
	massEm        float64
	radiusEm      float64
	gravityG      float64
	temperatureK  float64
	rings         int
	wideRing      bool
	atmosphere    string
	landable      bool
	terraformable bool
	possibleBio   []string
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
