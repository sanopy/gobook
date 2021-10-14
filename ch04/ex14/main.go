// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/sanopy/gobook/ch04/ex14/github"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.Items | len}} issues</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>Issue Title</th>
	<th>State</th>
	<th>User</th>
	<th>Milestone Title</th>
	<th>Milestone State</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	{{ if .Milestone }}
		<td><a href='{{.Milestone.Url}}'>{{.Milestone.Title}}</a></td>
		<td>{{.Milestone.State}}</td>
	{{ else }}
		<td></td>
		<td></td>
	{{ end }}
</tr>
{{end}}
</table>
`))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		repo := r.FormValue("repo")

		result, err := github.GetIssues(repo)
		if err != nil {
			fmt.Fprintf(w, "%q", err)
			return
		}

		if err := issueList.Execute(w, result); err != nil {
			fmt.Fprintf(w, "%q", err)
			return
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
