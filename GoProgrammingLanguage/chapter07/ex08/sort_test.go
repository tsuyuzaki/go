package main

import (
	"./mytrack"
	"sort"
	"testing"
)

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
