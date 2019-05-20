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
	sortKeys []string
}

func NewTracks(tracks []*Track, keys []string) *TracksToBeSorted {
	return &TracksToBeSorted{Tracks: tracks, sortKeys: keys}
}

func (ts *TracksToBeSorted) Len() int {
	return len(ts.Tracks)
}

func (ts *TracksToBeSorted) Less(i, j int) bool {
	for _, key := range ts.sortKeys {
		if key == "" {
			key = "Title"
		}
		result := ts.compare(i, j, key)
		if result != 0 {
			return (result < 0)
		}
	}
	return (i < j)
}

func (ts *TracksToBeSorted) Swap(i, j int) {
	ts.Tracks[i], ts.Tracks[j] = ts.Tracks[j], ts.Tracks[i]
}

func (ts *TracksToBeSorted) compare(i, j int, key string) int {
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


func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}