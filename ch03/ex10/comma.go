package comma

import (
	"bytes"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaBuf(s string) string {
	var buf bytes.Buffer

	if len(s) <= 3 {
		return s
	}

	num := len(s) % 3
	if num == 0 {
		num = 3
	}
	buf.WriteString(s[:num])

	for i := num; i+3 <= len(s); i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}
