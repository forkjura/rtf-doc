// Package rtf-doc provides simple tools for creation and writing rtf documents.
// It is very early in development and has suck features as work with text
// (color, font, aligning), tables (merged cells, borders style, thickness and colors),
// and pictures (jpeg or png format)
package rtfdoc

import (
	"fmt"
	"image/color"
	"strings"
)

// default compose

// NewDocument returns new rtf Document instance
func NewDocument() *Document {
	doc := Document{
		orientation: OrientationPortrait,
		header:      getDefaultHeader(),
		content:     nil,
	}
	doc.marginLeft = 720
	doc.marginRight = 720
	doc.marginTop = 720
	doc.marginBottom = 720

	doc.SetFormat(FormatA4)
	doc.SetOrientation(OrientationPortrait)

	doc.maxWidth = doc.pagesize.width - doc.marginRight - doc.marginLeft

	// Default fonts
	ft := doc.NewFontTable()
	ft.AddFont("roman", 0, 2, "Times New Roman", FontTimesNewRoman)
	ft.AddFont("roman", 2, 2, "Symbol", FontSymbol)
	ft.AddFont("swiss", 0, 2, "Arial", FontArial)
	ft.AddFont("swiss", 0, 2, "Comic Sans MS", FontComicSansMS)
	ft.AddFont("modern", 128, 1, "Curier New", FontCourierNew)

	// Default colortable
	ct := doc.NewColorTable()
	ct.AddColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}, ColorBlack)
	ct.AddColor(color.RGBA{R: 0, G: 0, B: 255, A: 255}, ColorBlue)
	ct.AddColor(color.RGBA{R: 0, G: 255, B: 255, A: 255}, ColorAqua)
	ct.AddColor(color.RGBA{R: 0, G: 255, B: 0, A: 255}, ColorLime)
	ct.AddColor(color.RGBA{R: 0, G: 128, B: 0, A: 255}, ColorGreen)
	ct.AddColor(color.RGBA{R: 255, G: 0, B: 255, A: 255}, ColorMagenta)
	ct.AddColor(color.RGBA{R: 255, G: 0, B: 0, A: 255}, ColorRed)
	ct.AddColor(color.RGBA{R: 255, G: 255, B: 0, A: 255}, ColorYellow)
	ct.AddColor(color.RGBA{R: 255, G: 255, B: 255, A: 255}, ColorWhite)
	ct.AddColor(color.RGBA{R: 0, G: 0, B: 128, A: 255}, ColorNavy)
	ct.AddColor(color.RGBA{R: 0, G: 128, B: 128, A: 255}, ColorTeal)
	ct.AddColor(color.RGBA{R: 128, G: 0, B: 128, A: 255}, ColorPurple)
	ct.AddColor(color.RGBA{R: 128, G: 0, B: 0, A: 255}, ColorMaroon)
	ct.AddColor(color.RGBA{R: 128, G: 128, B: 0, A: 255}, ColorOlive)
	ct.AddColor(color.RGBA{R: 128, G: 128, B: 128, A: 255}, ColorGray)
	ct.AddColor(color.RGBA{R: 192, G: 192, B: 192, A: 255}, ColorSilver)

	return &doc
}

func (doc *Document) getMargins() string {
	return fmt.Sprintf("\n\\margl%d\\margr%d\\margt%d\\margb%d",
		doc.marginLeft,
		doc.marginRight,
		doc.marginTop,
		doc.marginBottom)
}

func (doc *Document) compose() string {
	var result strings.Builder
	result.WriteString("{")
	result.WriteString(doc.header.compose())
	if doc.orientation == OrientationLandscape {
		result.WriteString(fmt.Sprintf("\n\\landscape"))
	}
	if doc.pagesize != (size{}) {
		result.WriteString(fmt.Sprintf("\n\\paperw%d\\paperh%d", doc.pagesize.width, doc.pagesize.height))
	}

	result.WriteString(doc.getMargins())

	for _, c := range doc.content {
		result.WriteString(fmt.Sprintf("\n%s", c.compose()))
	}
	result.WriteString("\n}")
	return result.String()
}

// SetFormat sets page format (A2, A3, A4)
func (doc *Document) SetFormat(format string) *Document {
	doc.pageFormat = format
	if doc.orientation != "" {
		size, err := getSize(format, doc.orientation)
		if err == nil {
			doc.pagesize = size
		}
	}
	doc.updateMaxWidth()

	return doc
}

// SetOrientation - sets page orientation (portrait, landscape)
func (doc *Document) SetOrientation(orientation string) *Document {

	for _, i := range []string{OrientationLandscape, OrientationPortrait} {
		if orientation == i {
			doc.orientation = i
		}
	}
	size, err := getSize(doc.pageFormat, doc.orientation)
	if err == nil {
		doc.pagesize = size
	}
	doc.updateMaxWidth()

	return doc
}

// // GetDocumentWidth - returns Document width
// func (doc *Document) GetDocumentWidth() int {
// 	return doc.pagesize.width
// }

// SetMarginLeft sets left margin for Document work area
func (doc *Document) SetMarginLeft(value int) *Document {
	doc.marginLeft = value
	return doc
}

// SetMarginRight sets right margin for Document work area
func (doc *Document) SetMarginRight(value int) *Document {
	doc.marginRight = value
	return doc
}

// SetMarginTop sets top margin for Document work area
func (doc *Document) SetMarginTop(value int) *Document {
	doc.marginTop = value
	return doc
}

// SetMarginBottom sets bottom margin for Document work area
func (doc *Document) SetMarginBottom(value int) *Document {
	doc.marginBottom = value
	return doc
}

// NewColorTable returns new color table for Document
func (doc *Document) NewColorTable() *ColorTable {
	ct := ColorTable{}
	doc.header.colorTable = &ct
	return &ct
}

// NewFontTable returns new font table for Document
func (doc *Document) NewFontTable() *FontTable {
	ft := FontTable{}
	doc.header.fontColor = &ft
	return &ft
}

func (doc *Document) getMaxWidth() int {
	return doc.maxWidth
}

func (doc *Document) GetMaxContentWidth() int {
	return doc.getMaxWidth()
}

func (doc *Document) updateMaxWidth() {
	doc.maxWidth = doc.pagesize.width - doc.marginRight - doc.marginLeft
}

// GetTableCellWidthByRatio - returns slice of cell widths from cells ratios
func (doc *Document) GetTableCellWidthByRatio(tableWidth int, ratio ...float64) []int {
	tw := tableWidth
	if tw > doc.maxWidth {
		tw = doc.maxWidth
	}
	cellRatioSum := 0.0
	for _, cellRatio := range ratio {
		cellRatioSum += cellRatio
	}
	var cellWidth = make([]int, len(ratio))
	for i := range ratio {
		cellWidth[i] = int(ratio[i] * (float64(tw) / cellRatioSum))
	}
	return cellWidth
}

// Export exports Document
func (doc *Document) Export() []byte {
	return []byte(doc.compose())
}

// AddFont function adds font to Document header
func (doc *Document) AddFont(family string, charset int, prq int, name string, code string) *Document {
	doc.fontColor.AddFont(family, charset, prq, name, code)
	return doc
}

// AddColor function adds colot to Document color table
func (doc *Document) AddColor(c color.RGBA, name string) *Document {
	doc.colorTable.AddColor(c, name)
	return doc
}
