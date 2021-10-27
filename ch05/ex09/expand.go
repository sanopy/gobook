package expand

import (
	"regexp"
)

var r = regexp.MustCompile(`\$\S*`)

func expand(s string, f func(string) string) string {
	if f == nil {
		return s
	}
	return r.ReplaceAllStringFunc(s, func(p string) string {
		return f(p[1:])
	})
}
