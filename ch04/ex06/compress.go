package compress

import (
	"unicode"
	"unicode/utf8"
)

func compressSpaces(b []byte) []byte {
	if len(b) <= 1 {
		return b
	}

	isContinuousSpace := false
	cnt := 0
	for i := 0; i < len(b); {
		r, s := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(r) {
			if !isContinuousSpace {
				b[cnt] = byte(' ')
				cnt++
			}
			isContinuousSpace = true
		} else {
			copy(b[cnt:cnt+s], b[i:i+s])
			cnt += s
			isContinuousSpace = false
		}
		i += s
	}
	return b[:cnt]
}
