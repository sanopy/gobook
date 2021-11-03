package mystrings

func Join(sep string, elms ...string) string {
	if len(elms) == 0 {
		return ""
	}

	s := elms[0]
	for _, elm := range elms[1:] {
		s += sep + elm
	}
	return s
}
