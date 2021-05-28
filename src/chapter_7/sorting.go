package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func main() {
	// using *track because swapping will be faster with pointers
	tracks := []*track{
		{
			Title:  "Go",
			Artist: "Delilah",
			Album:  "From the Roots Up",
			Year:   2012,
			Length: length("3m38s"),
		},
		{
			Title:  "Go",
			Artist: "Moby",
			Album:  "Moby",
			Year:   1992,
			Length: length("3m37s"),
		},
		{
			Title:  "Go Ahead",
			Artist: "Alicia Keys",
			Album:  "As I Am",
			Year:   2007,
			Length: length("4m36s"),
		},
		{
			Title:  "Ready 2 Go",
			Artist: "Martin Solveig",
			Album:  "Smash",
			Year:   2011,
			Length: length("4m24s"),
		},
	}

	sort.Sort(byArtist(tracks))

	printTracks(tracks)
}

func length(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		log.Panic(err)
	}

	return duration
}

func printTracks(tracks []*track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")

	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")

	for _, track := range tracks {
		fmt.Fprintf(tw, format, track.Title, track.Artist, track.Album, track.Year, track.Length)
	}

	tw.Flush()
}

// All this work instead of using a higher order function, yikes.
// tracks := []track{...}
// sortedTracks := tracks.Sort(func (trackA, trackB track) bool { return trackA.Artist < trackB.Artist })

type byArtist []*track

func (tracks byArtist) Len() int {
	return len(tracks)
}

func (tracks byArtist) Less(i, j int) bool {
	return tracks[i].Artist < tracks[j].Artist
}

func (tracks byArtist) Swap(i, j int) {
	tracks[i], tracks[j] = tracks[j], tracks[i]
}
