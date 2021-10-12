package rev

const LENGTH = 8

func reverse(s *[LENGTH]int) {
	for i, j := 0, LENGTH-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
