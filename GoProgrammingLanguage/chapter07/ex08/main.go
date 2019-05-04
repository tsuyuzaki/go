/**
 * 多くの GUI は、状態を持つ多段ソートの表ウィジェットを提供しています。
 * 一次ソートキーは最も直近にクリックされた列の見出し、二次ソートキーは二番目に近くクリックされた列の見出しといった具合になります。
 * このような表が使う sort.Interface の実装を定義しなさい。
 * その実装を sort.Stable を使う繰り返しソートと比較しなさい。
 *
 * > go test -bench . -benchmem sort_test.go
 */
package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"sort"
	"./mytrack"
)

var tracks = []*mytrack.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, mytrack.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, mytrack.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, mytrack.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, mytrack.Length("4m24s")},
}

func printTracks(tracks []*mytrack.Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func main() {
	fmt.Println("-----")
	ts := mytrack.TracksToBeSorted{tracks, "Title", "Album"}
	sort.Sort(ts)
	printTracks(ts.Tracks)

	fmt.Println("\n\n-----")
	ts.SecondlyKey = "Year"
	sort.Sort(ts)
	printTracks(ts.Tracks)

	fmt.Println("\n\n-----")
	ts.PrimaryKey = "Artist"
	sort.Sort(ts)
	printTracks(ts.Tracks)
}