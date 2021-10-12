package rotate

// https://leetcode.com/problems/rotate-array/discuss/54289/My-three-way-to-solve-this-problem-the-first-way-is-interesting(JAVA)
func rotate(s []int, k uint) {
	if len(s) <= 1 {
		return
	}

	//step each time to move
	step := int(k) % len(s)
	//find GCD between nums length and step
	gcd := gcd(len(s), step)

	//gcd path to finish movie
	for i := 0; i < gcd; i++ {
		//beginning position of each path
		position := i
		//count is the number we need swap each path
		count := len(s)/gcd - 1
		for j := 0; j < count; j++ {
			position = (position + step) % len(s)
			//swap index value in index i and position
			s[i], s[position] = s[position], s[i]
		}
	}
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
