package file

import (
	"fmt"

	"github.com/danwakefield/fnmatch"
)

// make fancy body name and color it
func CB(body string, f ...int) (ret string) {

	if body == "" {
		return "?"
	}

	var format string

	if f == nil {
		format = `<span weight="bold" color="%s">%s</span>`
	} else {
		if f[0] == -4 {
			format = `<span weight="bold" color="%s">%-4.4s</span>`
		} else if f[0] == 4 {
			format = `<span weight="bold" color="%s">%+4.4s</span>`
		}
		if f[0] == -5 {
			format = `<span weight="bold" color="%s">%-5.5s</span>`
		} else if f[0] == 5 {
			format = `<span weight="bold" color="%s">%5.5s</span>`
		}
	}

	color := "#EEEEEE"
	name := body

	defer func() {
		ret = fmt.Sprintf(format, color, name)
	}()

	// planets
	switch body {
	case "Earthlike body":
		name = "ELW"
		color = "#50FF50"
		return
	case "Water world":
		name = "WW"
		color = "#5050FF"
		return
	case "Ammonia world":
		name = "AW"
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
		color = "#FF9080"
		return
	case "Rocky ice body":
		name = "R/Ice"
		color = "#FFFFA0"
		return
	case "Gas giant with ammonia based life":
		name = "GG/AL"
		return
	case "Gas giant with water based life":
		name = "GG/WL"
		return
	case "Sudarsky class I gas giant":
		name = "GG 1c"
		return
	case "Sudarsky class II gas giant":
		name = "GG 2c"
		return
	case "Sudarsky class III gas giant":
		name = "GG 3c"
		return
	case "Sudarsky class IV gas giant":
		name = "GG 4c"
		return
	case "Sudarsky class V gas giant":
		name = "GG 5c"
		return
	case "Water giant":
		name = "WG"
		return
	case "Helium rich gas giant":
		name = "HeRGG"
		return
	}

	// rename stars like A_BlueWhiteSuperGiant or K_OrangeGiant
	if fnmatch.Match("*Giant*", body, 0) {
		name = string(body[0]) + "!"
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
		name = "@"
		color = "#A0A0A0"
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
