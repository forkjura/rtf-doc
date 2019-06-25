package rtfdoc

import (
	"fmt"
	"strings"
)

// AddParagraph return new instance of Paragraph
func (doc *Document) AddParagraph() *Paragraph {
	p := Paragraph{
		align:   AlignCenter,
		indent:  "\\fl360",
		content: nil,
		generalSettings: generalSettings{
			colorTable: doc.colorTable,
			fontColor:  doc.fontColor,
		},
		allowedWidth: doc.maxWidth,
	}
	p.updateMaxWidth()
	doc.content = append(doc.content, &p)
	return &p
}

func (par *Paragraph) updateMaxWidth() *Paragraph {
	par.maxWidth = par.allowedWidth
	return par
}

func (par Paragraph) compose() string {
	var res strings.Builder
	indentStr := fmt.Sprintf("\\fi%d \\li%d \\ri%d",
		par.indentFirstLine,
		par.indentLeftIndent,
		par.indentRightIndent)
	res.WriteString(fmt.Sprintf("\n\\pard \\q%s {%s ", par.align, indentStr))
	if par.isTable {
		res.WriteString("\\intbl")
	}
	// res += fmt.Sprintf(" \\q%s", par.align)

	for _, c := range par.content {
		res.WriteString(c.compose())
	}
	// res += "\n\\par}"
	res.WriteString("}")
	if !par.isTable {
		res.WriteString("\\par")
	}
	return res.String()
}

// SetIndentFirstLine function sets first line indent in twips
func (par *Paragraph) SetIndentFirstLine(value int) *Paragraph {
	par.indentFirstLine = value
	return par
}

// SetIndentRight function sets right indent in twips
func (par *Paragraph) SetIndentRight(value int) *Paragraph {
	par.indentRightIndent = value
	return par
}

// SetIndentLeft function sets left indent in twips
func (par *Paragraph) SetIndentLeft(value int) *Paragraph {
	par.indentLeftIndent = value
	return par
}

// SetAlign sets Paragraph align (c/center, l/left, r/right, j/justify)
func (par *Paragraph) SetAlign(align string) *Paragraph {
	for _, i := range []string{
		AlignCenter,
		AlignLeft,
		AlignRight,
		AlignJustify,
		AlignDistribute,
	} {
		if i == align {
			par.align = i
		}
	}

	return par
}
