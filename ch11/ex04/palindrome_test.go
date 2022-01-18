package palindrome

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // 24までのランダムな長さ
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // '\u0999' までのランダムなルーン
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomPalindromeIncludesSign(rng *rand.Rand) string {
	var s []rune
	for { // 0文字の場合は再生成
		s = []rune(randomPalindrome(rng))
		if len(s) > 0 {
			break
		}
	}
	n := rng.Intn(25) // 24までのランダムな長さ
	signs := []rune{' ', ',', '.', '\'', '"'}
	for i := 0; i < n; i++ {
		sign := signs[rng.Intn(len(signs))]
		pos := rng.Intn(len(s))
		// pos番目に要素を追加
		s = append(s[:pos+1], s[pos:]...)
		s[pos] = sign
	}
	return string(s)
}

func randomNotPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 2 // 2から26までのランダムな長さ
	runes := make([]rune, n)
	letterFlg := false
	for i := 0; i < (n+1)/2; i++ {
		var r1, r2 rune
		for {
			r1 = rune(rng.Intn(0x1000)) // '\u0999' までのランダムなルーン
			r2 = rune(rng.Intn(0x1000)) // '\u0999' までのランダムなルーン

			if unicode.IsLetter(r1) && unicode.IsLetter(r2) &&
				unicode.ToLower(r1) != unicode.ToLower(r2) {
				if i != n-1-i {
					letterFlg = true
				}
				break
			}
			if !unicode.IsLetter(r1) && !unicode.IsLetter(r2) {
				break
			}
		}
		runes[i] = r1
		runes[n-1-i] = r2
	}
	if !letterFlg {
		return randomNotPalindrome(rng) // 非文字のみの場合は再生成
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// 疑似乱数生成器を初期化する
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}

	for i := 0; i < 1000; i++ {
		p := randomPalindromeIncludesSign(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}

	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
