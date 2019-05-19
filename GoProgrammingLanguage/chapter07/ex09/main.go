/**
 * html/template パッケージ (4.6節) を使って、printTracks を HTML の表として曲を表示する関数で置き換えなさい。
 * 列の見出しをクリックしたら表をソートするために HTTP リクエストを行うように前の練習問題の解答を使いなさい*2。
 * *2訳注: この練習問題は、ウェブサーバを作成することを意図しています。
 */
package main

import (
	"fmt"
	"os"
	"log"
	"sort"
	"./mytrack"
	"html/template"
	"net/http"
)

var tracks = []*mytrack.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, mytrack.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, mytrack.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, mytrack.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, mytrack.Length("4m24s")},
}

func getSortKey(r *http.Request) (string, bool) {
	query := r.URL.Query()
	value, ok := query["sortkey"]
	if !ok {
		return "", false
	}
	if len(value) == 0 {
		return "", false
	}
	return value[0], true
}

func main() {
	table := template.Must(template.ParseFiles("./html/index.html"))
	tracks := mytrack.NewTracks(tracks)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key, ok := getSortKey(r)
		if ok {
			tracks.SetSortKey(key)
		}
		
		sort.Sort(tracks)
		
		if err := table.Execute(w, tracks); err != nil {
			fmt.Fprintf(os.Stderr, "table.Execute error [%v]\n", err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
