package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x byte) int {
	return int(pc[x])
}

func PopCountDiff(x, y [32]byte) int {
	sum := 0
	for i := 0; i < 32; i++ {
		sum += PopCount(x[i] ^ y[i])
	}
	return sum
}
