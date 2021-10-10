package main

import "fmt"

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)

func main() {
	fmt.Printf("KB: %dB\n", KB)
	fmt.Printf("MB: %dB\n", MB)
	fmt.Printf("GB: %dB\n", GB)
	fmt.Printf("TB: %dB\n", TB)
	fmt.Printf("PB: %dB\n", PB)
	fmt.Printf("EB: %dB\n", EB)
	// fmt.Printf("ZB: %dB\n", ZB)
	// fmt.Printf("YB: %dB\n", YB)
}
