package mytrack

import (
	"fmt"
	"os"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var validSortKeys = map[string]bool{"Title": true, "Artist": true, "Album": true, "Year": true, "Length": true}

type TracksToBeSorted struct {
	Tracks   []*Track
	sortKeys []string
}

func NewTracks(tracks []*Track, keys []string) *TracksToBeSorted {
	var sortKeys []string
	for _, key := range keys { // key validation
		if _, ok := validSortKeys[key]; ok {
			sortKeys = append(sortKeys, key)
		}
	}
	if len(sortKeys) == 0 { // For default order
		sortKeys = append(sortKeys, "Title")
	}
	return &TracksToBeSorted{Tracks: tracks, sortKeys: sortKeys}
}

func (ts *TracksToBeSorted) Len() int {
	return len(ts.Tracks)
}

func (ts *TracksToBeSorted) Less(i, j int) bool {
	for _, key := range ts.sortKeys {
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

func (ts *TracksToBeSorted) Swap(i, j int) {
	ts.Tracks[i], ts.Tracks[j] = ts.Tracks[j], ts.Tracks[i]
}

func (ts *TracksToBeSorted) less(i, j int, key string) bool {
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

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
