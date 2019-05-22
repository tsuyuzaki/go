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

func getSortKeys(r *http.Request) []string {
	key := r.URL.Query().Get(sortKeyName)
	if key == "" {
		return []string{} // For default order
	}
	var keys []string
	keys = append(keys, key)
	cookie, err := r.Cookie(sortKeyName)
	if err == nil && cookie.Value != "" {
		keys = append(keys, cookie.Value)
	}
	return keys
}

func setSortKeyToCookie(w http.ResponseWriter, key string) {
	cookie := &http.Cookie{Name: sortKeyName, Value: key}
	http.SetCookie(w, cookie)
}

func main() {
	table := template.Must(template.ParseFiles("./html/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		keys := getSortKeys(r)
		if len(keys) != 0 {
			setSortKeyToCookie(w, keys[0])
		} else {
			setSortKeyToCookie(w, "") // Clear cookie
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
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
