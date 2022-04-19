package fwt

import (
	"fmt"
	"testing"
)

func TestText(t *testing.T) {

	tbl := &Table{
		Delimiter: " | ",
		Default:   "?????",
		Pango:     false,
	}

	tbl.Header(&Header{
		Text:    "Cell 1!!!!!!!!!!!",
		FgColor: "blue",
		BgColor: "red",
	})

	tbl.Header(&Header{
		Text: "     Yet One Cell",
	})

	tbl.Header(&Header{
		Text: "     Yet One Cell",
	})

	tbl.Header(&Header{
		Text: "     Yet One Cell",
	})

	tbl.Cell(0, &Cell{
		Text: "some text here",
		Bold: true,
		Left: true,
	})

	tbl.Cell(0, &Cell{
		Text:    "another text there",
		FgColor: "red",
		Left:    false,
	})

	tbl.Cell(3, &Cell{
		Text:    "YAY!",
		BgColor: "green",
		Left:    false,
	})

	tbl.Cell(2, &Cell{
		Text:    "very long line that need to be cut",
		FgColor: "green",
		Left:    false,
		Bold:    true,
		Italic:  true,
	})

	tbl.Cell(2, &Cell{
		Text:    "very long line that need to be cut",
		FgColor: "green",
		Left:    true,
	})

	tbl.Cell(2, &Cell{
		Text:      "very long line that need to be cut",
		BgColor:   "green",
		Left:      true,
		Underline: true,
	})

	fmt.Println(tbl.Text())

}
