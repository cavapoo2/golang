// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func sortByMostClicked(top string) {
	sort.Sort(customSort{tracks, func(x, y *Track) bool {

		switch top {
		case "Title":

			if x.Title != y.Title {
				return x.Title < y.Title
			}
		case "Year":

			if x.Year != y.Year {
				return x.Year < y.Year
			}
		case "Length":
			if x.Length != y.Length {
				return x.Length < y.Length
			}
		case "Artist":
			if x.Artist != y.Artist {
				return x.Artist < y.Artist
			}
		}
		return false
	}})

}

func shuffleTracks(choice []string) {
	rand.Seed(time.Now().UTC().UnixNano())
	//	dest := make([]string, len(clicks))
	//	perm := rand.Perm(len(clicks))
	/*
		for i, v := range perm {
			dest[v] = clicks[i]
		}*/
	for i := range choice {
		j := rand.Intn(len(choice))
		choice[i], choice[j] = choice[j], choice[i]
	}
	return

}

func main() {
	printTracks(tracks)
	clicks := []string{"Title", "Year", "Length", "Artist"}
	//shuffle
	shuffleTracks(clicks)
	println("")
	println("Sort by ", clicks[0])
	sortByMostClicked(clicks[0])
	printTracks(tracks)

}

//!+customcode
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
