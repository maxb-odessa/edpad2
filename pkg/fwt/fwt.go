package fwt

import (
	"fmt"
	"strings"
)

// Fixed Width Table

type cellData struct {
	Text      string
	Left      bool
	NoFormat  bool
	FgColor   string
	BgColor   string
	Italic    bool
	Bold      bool
	Underline bool
}

type Header cellData
type Cell cellData

type Table struct {
	Delimiter string
	Default   string
	Pango     bool
	headers   []*Header
	cells     [][]*Cell
}

// add header to the table
func (t *Table) Header(h *Header) {

	// init headers array if not yet
	if len(t.headers) == 0 {
		t.headers = make([]*Header, 0)
	}

	// init cells 2d array if not yet
	if len(t.cells) == 0 {
		t.cells = make([][]*Cell, 0)
	}

	t.headers = append(t.headers, h)
}

// add cell to the table
func (t *Table) Cell(row int, c *Cell) {

	// no such row yet present
	if len(t.cells) <= row {

		// create new cell rows after already existing
		for i := len(t.cells); i <= row; i++ {
			// append new cells row to cells array
			t.cells = append(t.cells, make([]*Cell, 0))
		}

	}

	t.cells[row] = append(t.cells[row], c)

	// extra cell in the row, no header for this column defined:
	//  create empty anonymous header for this cell and mark the cell as 'do not format'
	if len(t.headers) < len(t.cells[row]) {
		t.headers = append(t.headers, &Header{})
		c.NoFormat = true
	}
}

// return formatted table
func (t *Table) Text() string {

	var text string

	// make header
	for _, h := range t.headers {

		elem := h.Text

		if t.Pango {
			c := *h
			elem = pangofyText(cellData(c), elem)
		}

		text += t.Delimiter + elem

	}

	text += t.Delimiter + "\n"

	// add rows
	for _, row := range t.cells {

		// add columns
		for idx := 0; idx < len(t.headers); idx++ {

			// fake cell with empty one if missed
			c := &Cell{Text: t.Default}
			if len(row) > idx {
				c = row[idx]
			}

			// formating of this cell is disabled
			cellElem := ""
			if c.NoFormat {
				cellElem = c.Text
			} else {

				// cut cell text to match header len
				cellSize := len([]rune(t.headers[idx].Text))
				cellText := c.Text
				if len([]rune(c.Text)) > cellSize {
					if c.Left {
						cellText = cellText[(len([]rune(cellText)))-cellSize:]
					} else {
						cellText = cellText[0:cellSize]
					}
				}

				format := "%*s"
				if c.Left {
					format = "%-*s"
				}

				cellElem = fmt.Sprintf(format, cellSize, cellText)
			}

			// apply pango style if requested
			if t.Pango {
				cd := *c
				cellElem = pangofyText(cellData(cd), cellElem)
			}

			text += t.Delimiter + cellElem
		}

		text += t.Delimiter + "\n"
	}

	// done
	return text
}

func pangofyText(c cellData, text string) string {

	strings.ReplaceAll(text, `<`, `&lt;`)
	strings.ReplaceAll(text, `>`, `&gt;`)
	strings.ReplaceAll(text, `&`, `&amp;`)

	if c.FgColor != "" || c.BgColor != "" {
		clr := ""
		if c.FgColor != "" {
			clr = ` fgcolor="` + c.FgColor + `"`
		}
		if c.BgColor != "" {
			clr += ` bgcolor="` + c.BgColor + `"`
		}
		text = `<span` + clr + `>` + text + `</span>`
	}

	if c.Bold {
		text = `<b>` + text + `</b>`
	}

	if c.Underline {
		text = `<u>` + text + `</u>`
	}

	if c.Italic {
		text = `<i>` + text + `</i>`
	}

	return text
}
