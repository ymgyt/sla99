package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

var (
	max = flag.Int("9num", 5, "number of nine")
)

type entry struct {
	rate   float64
	hour   float64
	minute float64
	second float64
}

func getEntries() []entry {
	year := 1.0 * 60 * 60 * 24 * 365
	nineNum := *max + 2

	var entiries = make([]entry, 0, *max-1)
	for i := 1; i < nineNum; i++ {
		v := math.Pow(10, float64((i + 1)))
		v = v - 1
		upRate := float64(v) * (1 / math.Pow(10, float64((i+1))))
		downRate := 1.0 - upRate

		downTimeSec := year * downRate
		entiries = append(entiries, entry{
			rate:   upRate,
			hour:   downTimeSec / 60 / 60,
			minute: downTimeSec / 60,
			second: downTimeSec,
		})
	}
	return entiries
}

func header() []string {
	return []string{
		"rate", "hour", "minute", "second",
	}
}

func showEntries(entries []entry) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Fprintln(w, strings.Join(header(), "\t"))
	for _, e := range entries {
		fmt.Fprintf(w, "%."+strconv.Itoa(*max)+"f\t%f\t%f\t%f\n", e.rate*100, e.hour, e.minute, e.second)
	}
	w.Flush()
}

func run(args []string) int {
	flag.Parse()
	entries := getEntries()
	showEntries(entries)
	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
