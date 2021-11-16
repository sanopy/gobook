package palindrome

import (
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	i := 0
	j := s.Len() - 1
	for i < j {
		if !s.Less(i, j) && !s.Less(j, i) { // s[i] == s[j]
			i++
			j--
		} else { // s[i] != s[j]
			return false
		}
	}
	return true
}
