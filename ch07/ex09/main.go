package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"text/template"
)

var trackList = template.Must(template.New("tracklist").Parse(`
<h1>{{.Tracks | len}} tracks</h1>
<table>
<tr style='text-align: left'>
	<th><a href="?sort=title">Title</a></th>
	<th><a href="?sort=artist">Artist</a></th>
	<th><a href="?sort=album">Album</a></th>
	<th><a href="?sort=year">Year</a></th>
	<th><a href="?sort=length">Length</a></th>
</tr>
{{range .Tracks}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

func main() {
	mks := NewMultiKeySort(tracks)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.FormValue("sort") {
		case "title":
			mks.AddCompare(TitleCompare)
		case "artist":
			mks.AddCompare(ArtistCompare)
		case "album":
			mks.AddCompare(AlbumCompare)
		case "year":
			mks.AddCompare(YearCompare)
		case "length":
			mks.AddCompare(LengthCompare)
		}

		sort.Sort(mks)
		if err := trackList.Execute(w, Tracks{tracks}); err != nil {
			fmt.Fprintf(w, "%q", err)
			return
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
