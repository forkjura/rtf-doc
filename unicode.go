package rtfdoc

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf16"
)

func convertNonASCIIToUTF16(text string) string {
	var res strings.Builder
	for _, r := range text {
		if r <= kEndOfASCII {
			res.WriteString(string(r))
		} else if r <= kMaxSigned16BitValue {
			//unicode plane 0 (values 0 - 32767)
			res.WriteString(fmt.Sprintf("\\u%d\\'5f", r))
		} else if r < kStartOfUnicodePlane1 {
			//unicode plane 0 (values 32768 - 65535)
			res.WriteString(fmt.Sprintf("\\u%d\\'5f", r-kUnsigned16BitValueIntoSigned16BitValueRange))
		} else {
			//unicode plane 1 upwards
			//convert unicode plane 1 to utf16
			r = r - kStartOfUnicodePlane1

			w1 := kUTF16HighSurrogateStart + ((r & kHigh10Bits) / kHigh10BitsShiftToLow10Bits) - kUnsigned16BitValueIntoSigned16BitValueRange
			w2 := kUTF16LowSurrogateStart + (r & kLow10Bits) - kUnsigned16BitValueIntoSigned16BitValueRange

			//add the high & low surrogates
			res.WriteString(fmt.Sprintf("\\u%d\\'5f\\u%d\\'5f", w1, w2))
		}
	}
	return res.String()
}
