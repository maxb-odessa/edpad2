package display

// utf pictograms
const (
	BLACKHOLE   = "\u269D"
	STAR4FILLED = "\u2726"
	STAR4HOLLOW = "\u2727"

	CIRCLEDN = "\u24DD"
	CIRCLEDY = "\u24E8"

	DASHEDRIGHTARROW      = "\u279F"
	HEAVYDASHEDRIGHTARROW = "\u27A0"
)

func CIRCLEDNUM(num int) string {
	if num <= 0 || num > 20 {
		return string(rune(0x24EA))
	} else {
		return string(rune(0x2460 + num))
	}
}

func CIRCLEDCHAR(char byte) string {
	return string(rune(0x24B6 + int(char) - 1))

}
