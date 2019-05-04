/**
 * 多くの GUI は、状態を持つ多段ソートの表ウィジェットを提供しています。
 * 一次ソートキーは最も直近にクリックされた列の見出し、二次ソートキーは二番目に近くクリックされた列の見出しといった具合になります。
 * このような表が使う sort.Interface の実装を定義しなさい。
 * その実装を sort.Stable を使う繰り返しソートと比較しなさい。
 */
package mytrack

import (
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
	Tracks []*Track
	PrimaryKey string
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
	result := ts.compare(i, j, ts.PrimaryKey)
	if result != 0 {
		return (result < 0)
	}
	result = ts.compare(i, j, ts.SecondlyKey)
	if result != 0 {
		return (result < 0)
	}
	return (i < j)
}

func (ts TracksToBeSorted) Swap(i, j int) {
	ts.Tracks[i], ts.Tracks[j] = ts.Tracks[j], ts.Tracks[i]
}

func (ts TracksToBeSorted) compare(i, j int, key string) int {
	lhs := ts.Tracks[i]
	rhs := ts.Tracks[j]

	if key == "Title" {
		if lhs.Title == rhs.Title {
			return 0
		} else if lhs.Title > rhs.Title {
			return 1
		} else {
			return -1
		}
	} else if key == "Artist" {
		if lhs.Artist == rhs.Artist {
			return 0
		} else if lhs.Artist > rhs.Artist {
			return 1
		} else {
			return -1
		}
	} else if key == "Album" {
		if lhs.Album == rhs.Album {
			return 0
		} else if lhs.Album > rhs.Album {
			return 1
		} else {
			return -1
		}
	} else if key == "Year" {
		if lhs.Year == rhs.Year {
			return 0
		} else if lhs.Year > rhs.Year {
			return 1
		} else {
			return -1
		}
	} else if key == "Length" {
		if lhs.Length == rhs.Length {
			return 0
		} else if lhs.Length > rhs.Length {
			return 1
		} else {
			return -1
		}
	} else {
		return (i - j)
	}
}
