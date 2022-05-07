package file

import (
	"fmt"

	"github.com/danwakefield/fnmatch"
)

// make fancy body name and color it
func CB(body string) (name, color string) {

	if body == "" {
		return "?", "#AAAAAA"
	}

	color = "#EEEEEE"
	name = body

	// planets
	switch body {
	case "Earthlike body":
		name = "ELW"
		color = "#50FF50"
		return
	case "Water world":
		name = "Water"
		color = "#5050FF"
		return
	case "Ammonia world":
		name = "Ammon"
		color = "#FF5050"
		return
	case "High metal content body":
		name = "HMC"
		color = "#A0A0FF"
		return
	case "Metal rich body":
		name = "Metal"
		color = "#E0E0FF"
		return
	case "Icy body":
		name = "Icy"
		color = "#FFFFFF"
		return
	case "Rocky body":
		name = "Rocky"
		color = "#CD7F32"
		return
	case "Rocky ice body":
		name = "RoIce"
		color = "#FFE070"
		return
	case "Gas giant with ammonia based life":
		name = "GgABL"
		return
	case "Gas giant with water based life":
		name = "GgWBL"
		return
	case "Sudarsky class I gas giant":
		name = "GgI"
		return
	case "Sudarsky class II gas giant":
		name = "GgII"
		return
	case "Sudarsky class III gas giant":
		name = "GgIII"
		return
	case "Sudarsky class IV gas giant":
		name = "GgIV"
		return
	case "Sudarsky class V gas giant":
		name = "GgV"
		return
	case "Water giant":
		name = "WG"
		return
	case "Helium rich gas giant":
		name = "GgHeR"
		return
	case "Helium gas giant":
		name = "GgHe"
		return
	}

	// rename stars like A_BlueWhiteSuperGiant or K_OrangeGiant
	if fnmatch.Match("*Giant*", body, fnmatch.FNM_IGNORECASE) {
		name = string(body[0]) + "+"
	}

	// color stars
	switch name[0] {
	case 'O':
		color = "#FFFFFF"
	case 'B':
		color = "#E0E0FF"
	case 'A':
		if name == "AeBe" {
			color = "#FFFF80"
		} else {
			color = "#FFFFE0"
		}
	case 'F':
		color = "#FFFFA0"
	case 'G':
		color = "#EEEE30"
	case 'K':
		color = "#FF9020"
	case 'M':
		if name == "MS" {
			color = "#FF7010"
		} else {
			color = "#FF3030"
		}
	case 'C':
		color = "#AAAAAA"
	case 'D':
		color = "#EEEEFF"
	case 'H':
		color = "#909090"
	case 'L', 'Y', 'T':
		color = "#FF2000"
	case 'N':
		color = "#AAAAFF"
	case 'S':
		color = "#CCCC70"
	case 'W':
		color = "#E0E0E0"
	}

	return
}

// in meters
const (
	SOLAR_RADIUS = 696340000.0
	EARTH_RADIUS = 6371.0 * 1000.0
	LIGHT_SECOND = 299792.0 * 1000.
)

func formatLargeNum(val float64) string {
	if val >= 1000000.0 {
		return fmt.Sprintf("%3.1fM", val/1000000.0)
	} else if val >= 1000.0 {
		return fmt.Sprintf("%3.1fK", val/1000.0)
	} else {
		return fmt.Sprintf("%4.0f", val)
	}
}
