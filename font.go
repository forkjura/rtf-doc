package rtfdoc

import (
	"fmt"
	"strings"
)

// AddFont returns font instance
func (ft *FontTable) AddFont(family string, charset int, prq int, name string, code string) *FontTable {
	if prq == 0 {
		prq = 2
	}
	*ft = append(*ft, font{family: family, charset: charset, prq: prq, name: name, code: code})
	return ft
}

func (f *font) encode() string {
	return fmt.Sprintf("\\f%s\\fprq%d\\fcharset%d %s;", f.family, f.prq, f.charset, f.name)

}

// FontTable

// NewFontTable - returns new font table
func NewFontTable() *FontTable {
	return &FontTable{}
}

func (ft FontTable) encode() string {
	var fontInfo strings.Builder
	for i := range ft {
		fontInfo.WriteString(fmt.Sprintf("{\\f%d%s}", i, ft[i].encode()))

	}
	return fontInfo.String()
}
