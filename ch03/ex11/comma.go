package comma

import "strings"

func comma(s string) string {
	if s == "" {
		return ""
	}

	if s[0] == '+' || s[0] == '-' {
		return string(s[0]) + comma(s[1:])
	}

	if idx := strings.Index(s, "."); idx != -1 {
		return comma(s[:idx]) + s[idx:]
	}

	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
