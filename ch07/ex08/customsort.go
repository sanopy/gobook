package main

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func TitleLess(x, y *Track) bool  { return x.Title < y.Title }
func ArtistLess(x, y *Track) bool { return x.Artist < y.Artist }
func AlbumLess(x, y *Track) bool  { return x.Album < y.Album }
func YearLess(x, y *Track) bool   { return x.Year < y.Year }
func LengthLess(x, y *Track) bool { return x.Length < y.Length }
