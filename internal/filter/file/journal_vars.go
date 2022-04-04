package file

var CurrentSystemName string
var CurrentMainStarName string
var CurrentMainStarClass string
var NextJumpSystem string

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
	shortName     string
	discovered    bool
	mapped        bool
	class         string
	massEm        float64
	radiusEm      float64
	gravityG      float64
	temperatureK  float64
	rings         int
	ringRad       float64
	atmosphere    string
	landable      bool
	terraformable bool
	signals       *bodySignals
}

type signalData struct {
	not_used_atm bool
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
