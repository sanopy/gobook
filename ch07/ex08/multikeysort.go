package main

type multiKeySort struct {
	t       []*Track
	compare []func(x, y *Track) int
}

func NewMultiKeySort(t []*Track) multiKeySort {
	m := multiKeySort{t, []func(x, y *Track) int{}}
	return m
}

func (x *multiKeySort) AddCompare(c func(x, y *Track) int) {
	x.compare = append(x.compare, c)
}

func (x *multiKeySort) T() []*Track {
	return x.t
}

func (x multiKeySort) Len() int { return len(x.t) }
func (x multiKeySort) Less(i, j int) bool {
	for k := len(x.compare) - 1; k >= 0; k-- {
		v := x.compare[k](x.t[i], x.t[j])
		if v != 0 {
			return v > 0
		}
	}
	return false
}
func (x multiKeySort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func TitleCompare(x, y *Track) int {
	if x.Title == y.Title {
		return 0
	} else if x.Title < y.Title {
		return 1
	} else {
		return -1
	}
}
func ArtistCompare(x, y *Track) int {
	if x.Artist == y.Artist {
		return 0
	} else if x.Artist < y.Artist {
		return 1
	} else {
		return -1
	}
}
func AlbumCompare(x, y *Track) int {
	if x.Album == y.Album {
		return 0
	} else if x.Album < y.Album {
		return 1
	} else {
		return -1
	}
}
func YearCompare(x, y *Track) int {
	if x.Year == y.Year {
		return 0
	} else if x.Year < y.Year {
		return 1
	} else {
		return -1
	}
}
func LengthCompare(x, y *Track) int {
	if x.Length == y.Length {
		return 0
	} else if x.Length < y.Length {
		return 1
	} else {
		return -1
	}
}
