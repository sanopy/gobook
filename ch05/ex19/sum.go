package sum

// sum of 0 to x
func sum(x int) (s int) {
	defer func() {
		if p := recover(); p != nil {
			s = p.(int)
		}
	}()

	if x == 0 {
		panic(0)
	}
	panic(x + sum(x-1))
}
