/**
 * 多くの GUI は、状態を持つ多段ソートの表ウィジェットを提供しています。
 * 一次ソートキーは最も直近にクリックされた列の見出し、二次ソートキーは二番目に近くクリックされた列の見出しといった具合になります。
 * このような表が使う sort.Interface の実装を定義しなさい。
 * その実装を sort.Stable を使う繰り返しソートと比較しなさい。
 */
package mytrack

import (
	"os"
	"fmt"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type TracksToBeSorted struct {
	Tracks      []*Track
	PrimaryKey  string
	SecondlyKey string
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func (ts TracksToBeSorted) Len() int {
	return len(ts.Tracks)
}

func (ts TracksToBeSorted) Less(i, j int) bool {
	keys := []string{ts.PrimaryKey, ts.SecondlyKey}
	for _, key := range keys {
		if ts.less(i, j, key) {
			return true
		}
		if !ts.less(j, i, key) { // equals
			continue
		}
		return false
	}
	return i < j
}

func (ts TracksToBeSorted) Swap(i, j int) {
	ts.Tracks[i], ts.Tracks[j] = ts.Tracks[j], ts.Tracks[i]
}

func (ts TracksToBeSorted) less(i, j int, key string) bool {
	lhs, rhs := ts.Tracks[i], ts.Tracks[j]

	if key == "Title" {
		return lhs.Title < rhs.Title
	} else if key == "Artist" {
		return lhs.Artist < rhs.Artist
	} else if key == "Album" {
		return lhs.Album < rhs.Album
	} else if key == "Year" {
		return lhs.Year < rhs.Year
	} else if key == "Length" {
		return lhs.Length < rhs.Length
	}
	fmt.Fprintf(os.Stderr, "Invalid key [%s].\n", key)
	return i < j
}
