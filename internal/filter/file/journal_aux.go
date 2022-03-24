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
		if f[0] == -3 {
			format = `<span weight="bold" color="%s">%-3.3s</span>`
		} else if f[0] == 3 {
			format = `<span weight="bold" color="%s">%+3.3s</span>`
		}
		if f[0] == -5 {
			format = `<span weight="bold" color="%s">%-5.5s</span>`
		} else if f[0] == 5 {
			format = `<span weight="bold" color="%s">%5.5s</span>`
		}
	}

	color := "#AAAAAA"
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
		color = "#A0A0E0"
		return
	case "Metal rich body":
		name = "Metal"
		color = "#E0E0F0"
		return
	case "Icy body":
		name = "Icy"
		color = "#FFFFFF"
		return
	case "Rocky body":
		name = "Rocky"
		color = "#D09080"
		return
	case "Rocky ice body":
		name = "R/Ice"
		color = "#F0F0A0"
		return
	case "Gas giant with ammonia based life",
		"Gas giant with water based life",
		"Sudarsky class I gas giant",
		"Sudarsky class II gas giant",
		"Sudarsky class III gas giant",
		"Sudarsky class IV gas giant",
		"Sudarsky class V gas giant",
		"Water giant":
		name = "GG"
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
		color = "#A00000"
	case 'N':
		color = "#AAAAFF"
	case 'S':
		color = "#AAAA70"
	case 'W':
		color = "#E0E0E0"
	}

	return
}
