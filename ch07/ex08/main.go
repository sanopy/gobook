package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Multi Key:")
	mks := NewMultiKeySort(tracks)
	mks.AddCompare(LengthCompare)
	mks.AddCompare(YearCompare)
	mks.AddCompare(TitleCompare)
	sort.Sort(mks)
	printTracks(tracks)

	fmt.Println("\nStable:")
	sort.Stable(customSort{tracks, LengthLess})
	sort.Stable(customSort{tracks, YearLess})
	sort.Stable(customSort{tracks, TitleLess})
	printTracks(tracks)
}
