package unique

func unique(s []string) []string {
	if len(s) <= 1 {
		return s
	}

	cnt := 1
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			s[cnt] = s[i]
			cnt++
		}
	}
	return s[:cnt]
}
