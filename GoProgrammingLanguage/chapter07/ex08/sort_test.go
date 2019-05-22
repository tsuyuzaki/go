package main

import (
	"./mytrack"
	"sort"
	"testing"
)

var tracks = []*mytrack.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, mytrack.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, mytrack.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, mytrack.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, mytrack.Length("4m24s")},
}

func BenchmarkMySort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ts := mytrack.TracksToBeSorted{tracks, "Title", "Album"}
		sort.Sort(ts)
	}
}

func BenchmarkStableSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ts := mytrack.TracksToBeSorted{tracks, "Title", "Album"}
		sort.Stable(ts)
	}
}
