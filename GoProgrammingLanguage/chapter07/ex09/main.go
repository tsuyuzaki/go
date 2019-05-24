/**
 * html/template パッケージ (4.6節) を使って、printTracks を HTML の表として曲を表示する関数で置き換えなさい。
 * 列の見出しをクリックしたら表をソートするために HTTP リクエストを行うように前の練習問題の解答を使いなさい*2。
 * *2訳注: この練習問題は、ウェブサーバを作成することを意図しています。
 */
package main

import (
	"./mytrack"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
)

const sortKeyName = "sortKey"
var table = template.Must(template.ParseFiles("./html/index.html"))

func getSortKeys(r *http.Request) []string {
	key := r.URL.Query().Get(sortKeyName)
	if key == "" {
		return []string{} // For default order
	}
    keys := []string{key}
	cookie, err := r.Cookie(sortKeyName)
	if err == nil && cookie.Value != "" {
		keys = append(keys, cookie.Value)
	}
	return keys
}

func setCookie(w http.ResponseWriter, name, value string) {
	cookie := &http.Cookie{Name: name, Value: value}
	http.SetCookie(w, cookie)
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	keys := getSortKeys(r)
	if len(keys) != 0 {
		setCookie(w, sortKeyName, keys[0])
	} else {
		setCookie(w, sortKeyName, "") // Clear cookie
	}
	tracks := []*mytrack.Track{
		{"Go", "Delilah", "From the Roots Up", 2012, mytrack.Length("3m38s")},
		{"Go", "Moby", "Moby", 1992, mytrack.Length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, mytrack.Length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, mytrack.Length("4m24s")},
	}
	ts := mytrack.NewTracks(tracks, keys)
	sort.Sort(ts)
	if err := table.Execute(w, ts.Tracks); err != nil {
		fmt.Fprintf(os.Stderr, "table.Execute error [%v]\n", err)
	}
}

func main() {
	http.HandleFunc("/", reqHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
