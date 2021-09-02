package rtfdoc

import (
	"fmt"
	"strings"
)

// SetDefaultFontSize sets default font size of Table
func (t *Table) SetDefaultFontSize(size int) *Table {
	t.defaultFontSize = size
	return t
}

// SetMarginLeft function sets Table left margin
func (t *Table) SetMarginLeft(value int) *Table {
	t.marginLeft = value
	return t
}

// SetMarginRight function sets Table right margin
func (t *Table) SetMarginRight(value int) *Table {
	t.marginRight = value
	return t
}

// SetMarginTop function sets Table top margin
func (t *Table) SetMarginTop(value int) *Table {
	t.marginTop = value
	return t
}

// SetMarginBottom function sets Table bottom margin
func (t *Table) SetMarginBottom(value int) *Table {
	t.marginBottom = value
	//tp.margins += fmt.Sprintf(" \\trpaddb%d", value)
	return t
}

// SetPaddingLeft function sets Table left margin
func (t *Table) SetPaddingLeft(value int) *Table {
	t.paddingLeft = value
	return t
}

// SetPaddingRight function sets Table right padding
func (t *Table) SetPaddingRight(value int) *Table {
	t.paddingRight = value
	return t
}

// SetPaddingTop function sets Table top padding
func (t *Table) SetPaddingTop(value int) *Table {
	t.paddingTop = value
	return t
}

// SetPaddingBottom function sets Table bottom padding
func (t *Table) SetPaddingBottom(value int) *Table {
	t.paddingBottom = value
	//tp.paddings += fmt.Sprintf(" \\trpaddb%d", value)
	return t
}

// SetPadding function sets all Table paddings
func (t *Table) SetPadding(value int) *Table {
	return t.SetPaddingBottom(value).SetPaddingLeft(value).SetPaddingRight(value).SetPaddingTop(value)
}

// SetAlign sets Table aligning (c/center, l/left, r/right)
func (t *Table) SetAlign(align string) *Table {
	for _, i := range []string{AlignCenter, AlignLeft, AlignRight} {
		if i == align {
			t.align = i
		}
	}
	return t
}

// AddTable returns Table instance
func (doc *Document) AddTable() *Table {
	t := Table{
		align:    AlignCenter,
		docWidth: doc.maxWidth,
	}
	t.SetMarginLeft(100).SetMarginRight(100).SetMarginTop(100).SetMarginBottom(100)

	t.colorTable = doc.colorTable
	t.fontColor = doc.fontColor
	t.SetBorderLeft(true).
		SetBorderRight(true).
		SetBorderTop(true).
		SetBorderBottom(true).
		SetBorderStyle(BorderSingleThickness).
		SetBorderColor(ColorBlack).
		SetBorderWidth(15)
	t.updateMaxWidth()
	doc.content = append(doc.content, &t)
	return &t
}

func (t *Table) updateMaxWidth() *Table {
	t.maxWidth = t.docWidth - t.marginLeft - t.marginRight
	return t
}

func (t Table) compose() string {
	var res strings.Builder
	var align = ""
	if t.align != "" {
		align = fmt.Sprintf("\\trq%s", t.align)
	}
	for _, tr := range t.data {
		res.WriteString(fmt.Sprintf("\n{\\trowd %s", align))
		if t.defaultFontSize > 0 {
			res.WriteString(fmt.Sprintf("\\fs%d", 2*t.defaultFontSize))
		}
		res.WriteString(fmt.Sprintf("\n\\trpaddl%d \\trpaddr%d \\trpaddt%d \\trpaddb%d\n", t.paddingLeft, t.paddingRight, t.paddingTop, t.paddingBottom))
		//res += t.getMargins()
		res.WriteString(tr.encode())
		res.WriteString("\\row}")
	}
	return res.String()
}

// AddTableRow returns new Table row instance
func (t *Table) AddTableRow() *TableRow {
	tr := TableRow{
		generalSettings: generalSettings{
			fontColor:  t.fontColor,
			colorTable: t.colorTable,
		},
		tableWidth: t.maxWidth,
	}
	tr.SetBorderLeft(t.borderLeft).
		SetBorderRight(t.borderRight).
		SetBorderTop(t.borderTop).
		SetBorderBottom(t.borderBottom).
		SetBorderStyle(t.borderStyle).
		SetBorderColor(t.borderColor).
		SetBorderWidth(t.borderWidth)
	t.updateMaxWidth()
	t.data = append(t.data, &tr)
	return &tr
}

func (tr *TableRow) updateMaxWidth() *TableRow {
	tr.maxWidth = tr.tableWidth
	return tr
}

// SetBorderLeft function sets Table left border presence
func (t *Table) SetBorderLeft(isBorder bool) *Table {
	t.borderLeft = isBorder
	return t
}

// SetBorderRight function sets Table right border presence
func (t *Table) SetBorderRight(isBorder bool) *Table {
	t.borderRight = isBorder
	return t
}

// SetBorderTop function sets Table top border presence
func (t *Table) SetBorderTop(isBorder bool) *Table {
	t.borderTop = isBorder
	return t
}

// SetBorderBottom function sets Table bottom border presence
func (t *Table) SetBorderBottom(isBorder bool) *Table {
	t.borderBottom = isBorder
	return t
}

// SetBorder function sets Table bottom border presence
func (t *Table) SetBorder(isBorder bool) *Table {
	t.borderBottom = isBorder
	return t.SetBorderBottom(isBorder).SetBorderTop(isBorder).SetBorderLeft(isBorder).SetBorderRight(isBorder)
}

// SetBorderStyle function sets Table left border style
func (t *Table) SetBorderStyle(bStyle string) *Table {
	for _, i := range []string{
		BorderDashSmall,
		BorderDashed,
		BorderDotDash,
		BorderDotDotDash,
		BorderDotted,
		BorderDouble,
		BorderDoubleThickness,
		BorderWavyDouble,
		BorderEmboss,
		BorderEngrave,
		BorderHairline,
		BorderInset,
		BorderOutset,
		BorderShadowed,
		BorderSingleThickness,
		BorderStripped,
		BorderThickThinLarge,
		BorderThickThinMedium,
		BorderThickThinSmall,
		BorderThinThickLarge,
		BorderThinThickMedium,
		BorderThinThickSmall,
		BorderThinThickThinLarge,
		BorderThinThickThinMedium,
		BorderTriple,
		BorderWavy,
	} {
		if bStyle == i {
			t.borderStyle = i
			for tr := range t.data {
				t.data[tr].SetBorderStyle(i)
			}
			break
		}
	}
	return t
}

// SetBorderColor function sets color of the Table's border and it's rows and cells
func (t *Table) SetBorderColor(color string) *Table {
	t.borderColor = color
	for tr := range t.data {
		t.data[tr].SetBorderColor(color)
	}
	return t
}

// SetBorderWidth function sets width of the Table's border and it's rows and cells
func (t *Table) SetBorderWidth(value int) *Table {
	t.borderWidth = value
	for tr := range t.data {
		t.data[tr].SetBorderWidth(value)
	}
	return t
}

// SetWidth sets width of Table
func (t *Table) SetWidth(width int) *Table {
	t.width = width
	return t
}

// SetBorderLeft function sets left border presence
func (tr *TableRow) SetBorderLeft(isBorder bool) *TableRow {
	tr.borderLeft = isBorder
	return tr
}

// SetBorderRight function sets right border presence
func (tr *TableRow) SetBorderRight(isBorder bool) *TableRow {
	tr.borderRight = isBorder
	return tr
}

// SetBorderTop function sets top border presence
func (tr *TableRow) SetBorderTop(isBorder bool) *TableRow {
	tr.borderTop = isBorder
	return tr
}

// SetBorderBottom function sets bottom border presence
func (tr *TableRow) SetBorderBottom(isBorder bool) *TableRow {
	tr.borderBottom = isBorder
	return tr
}

// SetBorder function sets bottom borders
func (tr *TableRow) SetBorder(isBorder bool) *TableRow {
	return tr.SetBorderBottom(isBorder).SetBorderTop(isBorder).SetBorderLeft(isBorder).SetBorderRight(isBorder)
}

// SetBorderStyle function sets border style
func (tr *TableRow) SetBorderStyle(bStyle string) *TableRow {
	for _, i := range []string{
		BorderDashSmall,
		BorderDashed,
		BorderDotDash,
		BorderDotDotDash,
		BorderDotted,
		BorderDouble,
		BorderDoubleThickness,
		BorderWavyDouble,
		BorderEmboss,
		BorderEngrave,
		BorderHairline,
		BorderInset,
		BorderOutset,
		BorderShadowed,
		BorderSingleThickness,
		BorderStripped,
		BorderThickThinLarge,
		BorderThickThinMedium,
		BorderThickThinSmall,
		BorderThinThickLarge,
		BorderThinThickMedium,
		BorderThinThickSmall,
		BorderThinThickThinLarge,
		BorderThinThickThinMedium,
		BorderTriple,
		BorderWavy,
	} {
		if bStyle == i {
			tr.borderStyle = i
			for c := range tr.cells {
				tr.cells[c].SetBorderStyle(i)
			}
			break
		}
	}
	return tr
}

// SetBorderColor sets border color of the row (and recursevely on its cells)
func (tr *TableRow) SetBorderColor(color string) *TableRow {
	tr.borderColor = color
	for c := range tr.cells {
		tr.cells[c].SetBorderColor(color)
	}
	return tr
}

// SetBorderWidth sets border width (and recursevely on its cells)
func (tr *TableRow) SetBorderWidth(value int) *TableRow {
	tr.borderWidth = value
	for c := range tr.cells {
		tr.cells[c].SetBorderWidth(value)
	}
	return tr
}

func (tr *TableRow) encode() string {
	var res strings.Builder
	// Border settings
	bTempl := "\n\\trbrdr%s\\brdrw%d\\brdr%s"
	for c := range *tr.colorTable {
		if ((*tr.colorTable)[c]).name == tr.borderColor {
			bTempl += fmt.Sprintf("\\brdrcf%d", c+1)
		}

	}

	if tr.borderLeft {
		res.WriteString(fmt.Sprintf(bTempl, "l", tr.borderWidth, tr.borderStyle))
	}
	if tr.borderRight {
		res.WriteString(fmt.Sprintf(bTempl, "r", tr.borderWidth, tr.borderStyle))
	}
	if tr.borderTop {
		res.WriteString(fmt.Sprintf(bTempl, "t", tr.borderWidth, tr.borderStyle))
	}
	if tr.borderBottom {
		res.WriteString(fmt.Sprintf(bTempl, "b", tr.borderWidth, tr.borderStyle))
	}

	if len(tr.cells) != 0 {
		cellLengthPosition := 0
		for _, tc := range tr.cells {

			cellLengthPosition += tc.getCellWidth()
			res.WriteString(tc.cellComposeProperties())
			res.WriteString(fmt.Sprintf("\\cellx%d", cellLengthPosition))

		}
		res.WriteString("\n")
		for _, tc := range tr.cells {
			res.WriteString(tc.cellComposeData())
		}
	}
	return res.String()
}

// AddDataCell returns new DataCell for current Table row
func (tr *TableRow) AddDataCell(width int) *TableCell {
	dc := TableCell{
		cellWidth: width,
		maxWidth:  width,
	}
	dc.fontColor = tr.fontColor
	dc.colorTable = tr.colorTable
	dc.SetBorderLeft(tr.borderLeft).
		SetBorderRight(tr.borderRight).
		SetBorderTop(tr.borderTop).
		SetBorderBottom(tr.borderBottom).
		SetBorderStyle(tr.borderStyle).
		SetBorderColor(tr.borderColor).
		SetBorderWidth(tr.borderWidth)
	dc.updateMaxWidth()
	tr.cells = append(tr.cells, &dc)
	return &dc
}

func (dc *TableCell) updateMaxWidth() *TableCell {
	dc.maxWidth = dc.cellWidth - dc.marginLeft - dc.marginRight
	return dc
}

// SetWidth sets width of the cell
func (dc *TableCell) SetWidth(cellWidth int) *TableCell {
	dc.cellWidth = cellWidth
	return dc
}

// AddParagraph creates cell's paragraph
func (dc *TableCell) AddParagraph() *Paragraph {
	p := Paragraph{
		isTable: true,
		align:   "l",
		indent:  "\\fl360",
		generalSettings: generalSettings{
			colorTable: dc.colorTable,
			fontColor:  dc.fontColor,
		},
		allowedWidth: dc.maxWidth,
	}
	p.updateMaxWidth()
	dc.content = append(dc.content, &p)
	return &p
}

func (dc TableCell) cellComposeProperties() string {
	var res strings.Builder
	// Тута свойства ячейки (границы, все дела...)
	bTempl := "\n\\clbrdr%s\\brdrw%d\\brdr%s"
	for c := range *dc.colorTable {
		if ((*dc.colorTable)[c]).name == dc.borderColor {
			bTempl += fmt.Sprintf("\\brdrcf%d", c+1)
		}

	}

	if dc.borderLeft {
		res.WriteString(fmt.Sprintf(bTempl, "l", dc.borderWidth, dc.borderStyle))
	}
	if dc.borderRight {
		res.WriteString(fmt.Sprintf(bTempl, "r", dc.borderWidth, dc.borderStyle))
	}
	if dc.borderTop {
		res.WriteString(fmt.Sprintf(bTempl, "t", dc.borderWidth, dc.borderStyle))
	}
	if dc.borderBottom {
		res.WriteString(fmt.Sprintf(bTempl, "b", dc.borderWidth, dc.borderStyle))
	}

	// Margins
	res.WriteString(fmt.Sprintf("\n\\clpadl%d\\clpadr%d\\clpadt%d\\clpadb%d",
		dc.paddingLeft, dc.paddingRight, dc.paddingTop, dc.paddingBottom,
	))

	// Vertical Merged
	if dc.verticalMerged != "" {
		res.WriteString(fmt.Sprintf("\\clvm%s", dc.verticalMerged))
	}

	// Aligning insite cell
	res.WriteString(fmt.Sprintf("\\clvertal%s", dc.vTextAlign))

	// Background Color

	if dc.backgroundColor != "" {
		for c := range *dc.colorTable {
			if ((*dc.colorTable)[c]).name == dc.backgroundColor {
				res.WriteString(fmt.Sprintf("\\clcbpat%d", c+1))
			}
		}
	}

	return res.String()
}

func (dc TableCell) cellComposeData() string {
	var res strings.Builder
	if len(dc.content) == 0 {
		dc.AddParagraph()
	}
	for _, p := range dc.content {
		res.WriteString(fmt.Sprintf("%s\n", p.compose()))
	}
	res.WriteString("\\cell")
	return res.String()
}

func (dc TableCell) getCellWidth() int {
	return dc.cellWidth
}

// SetBorders sets borders to
// datacell

// SetBorderLeft function set left border to be visible
func (dc *TableCell) SetBorderLeft(value bool) *TableCell {
	dc.borderLeft = value
	return dc
}

// SetBorderRight function sets right border to be visible
func (dc *TableCell) SetBorderRight(value bool) *TableCell {
	dc.borderRight = value
	return dc
}

// SetBorderTop function sets top border to be visible
func (dc *TableCell) SetBorderTop(value bool) *TableCell {
	dc.borderTop = value
	return dc
}

// SetBorderBottom function sets bottom border to be visible
func (dc *TableCell) SetBorderBottom(value bool) *TableCell {
	dc.borderBottom = value
	return dc
}

// SetBorder function sets bottom borders
func (dc *TableCell) SetBorder(isBorder bool) *TableCell {
	return dc.SetBorderBottom(isBorder).SetBorderTop(isBorder).SetBorderLeft(isBorder).SetBorderRight(isBorder)
}

// SetBorderWidth function sets cell's border width px
func (dc *TableCell) SetBorderWidth(value int) *TableCell {
	dc.borderWidth = value
	return dc
}

// SetBorderStyle function sets cell's border style
func (dc *TableCell) SetBorderStyle(bStyle string) *TableCell {
	bStyle = BorderSingleThickness
	for _, i := range []string{
		BorderDashSmall,
		BorderDashed,
		BorderDotDash,
		BorderDotDotDash,
		BorderDotted,
		BorderDouble,
		BorderDoubleThickness,
		BorderWavyDouble,
		BorderEmboss,
		BorderEngrave,
		BorderHairline,
		BorderInset,
		BorderOutset,
		BorderShadowed,
		BorderSingleThickness,
		BorderStripped,
		BorderThickThinLarge,
		BorderThickThinMedium,
		BorderThickThinSmall,
		BorderThinThickLarge,
		BorderThinThickMedium,
		BorderThinThickSmall,
		BorderThinThickThinLarge,
		BorderThinThickThinMedium,
		BorderTriple,
		BorderWavy,
	} {
		if bStyle == i {
			dc.borderStyle = i
			break
		}
	}
	return dc
}

// GetTableCellWidthByRatio returns slice of cell widths
func (t *Table) GetTableCellWidthByRatio(ratio ...float64) []int {

	cellRatioSum := 0.0
	for _, cellRatio := range ratio {
		cellRatioSum += cellRatio
	}
	var cellWidth = make([]int, len(ratio))
	for i := range ratio {
		cellWidth[i] = int(ratio[i] * (float64(t.width) / cellRatioSum))
	}
	return cellWidth
}

// SetVerticalMergedFirst sets this cell to be first in vertical merging.
func (dc *TableCell) SetVerticalMergedFirst() *TableCell {
	dc.verticalMerged = "gf"
	return dc
}

// SetVerticalMergedNext sets this cell to be not first cell in vertical merging.
func (dc *TableCell) SetVerticalMergedNext() *TableCell {
	dc.verticalMerged = "rg"
	return dc
}

// func (dc TableCell) getVerticalMergedProperty() string {
// 	return dc.verticalMerged
// }

// SetMarginLeft function sets this cell's left margin
func (dc *TableCell) SetMarginLeft(value int) *TableCell {
	dc.marginLeft = value
	return dc
}

// SetMarginRight function sets this cell's right margin
func (dc *TableCell) SetMarginRight(value int) *TableCell {
	dc.marginRight = value
	return dc
}

// SetMarginTop function sets this cell's top margin
func (dc *TableCell) SetMarginTop(value int) *TableCell {
	dc.marginTop = value
	return dc
}

// SetMarginBottom function sets this cell's bottom margin
func (dc *TableCell) SetMarginBottom(value int) *TableCell {
	dc.marginBottom = value
	return dc
}

// SetPaddingLeft function sets this cell's left padding
func (dc *TableCell) SetPaddingLeft(value int) *TableCell {
	dc.paddingLeft = value
	return dc
}

// SetPaddingRight function sets this cell's right padding
func (dc *TableCell) SetPaddingRight(value int) *TableCell {
	dc.paddingRight = value
	return dc
}

// SetPaddingTop function sets this cell's top padding
func (dc *TableCell) SetPaddingTop(value int) *TableCell {
	dc.paddingTop = value
	return dc
}

// SetPaddingBottom function sets this cell's bottom padding
func (dc *TableCell) SetPaddingBottom(value int) *TableCell {
	dc.paddingBottom = value
	return dc
}

// SetPadding - function sets all paddings to value
func (dc *TableCell) SetPadding(value int) *TableCell {
	return dc.SetPaddingBottom(value).SetPaddingLeft(value).SetPaddingRight(value).SetPaddingTop(value)
}

// SetVAlign sets align
func (dc *TableCell) SetVAlign(valign string) *TableCell {
	dc.vTextAlign = "t"
	for _, i := range []string{VAlignBottom, VAlignMiddle, VAlignTop} {
		if valign == i {
			dc.vTextAlign = i
		}
	}
	return dc
}

// SetBorderColor function sets cell's border color
func (dc *TableCell) SetBorderColor(color string) *TableCell {
	dc.borderColor = color
	return dc
}

// SetBackgroundColor function sets cell's background color
func (dc *TableCell) SetBackgroundColor(color string) *TableCell {
	dc.backgroundColor = color
	return dc
}
