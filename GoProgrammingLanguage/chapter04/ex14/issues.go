/**
 * GitHubへの一度の問い合わせで、バグレポート、マイルストーン、ユーザの一覧を閲覧可能にするウェブサーバを作りなさい。
 */
package main

import (
	"./github"
	"html/template"
	"log"
	"net/http"
)

var issuerList = template.Must(template.New("ussuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
    <th>#</th>
    <th>Statue</th>
    <th>User</th>
    <th>Title</th>
</tr>
{{range .Items}}
<tr>
    <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
    <td>{{.State}}</td>
    <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
    <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

// var args = []string{"repo:golang/go", "3133", "10535"}
var args = []string{"repo:golang/go", "commenter:gopherbot", "json", "encoder"}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		issues, err := github.SearchIssues(args)
		if err != nil {
			log.Fatal(err)
		}
		if err := issuerList.Execute(w, issues); err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
