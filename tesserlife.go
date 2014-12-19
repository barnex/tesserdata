// +build ignore

// Histogram of expected lifetimes
package main

import (
	"flag"
	"io"
	"log"
	"os"

	. "."
)

var (
	flag_res    = flag.Float64("res", 1, "Time resolution in seconds")
	flag_n      = flag.Int("bins", 100, "N bins")
	flag_player = flag.Bool("playerA", false, "Select player")
)

func main() {
	log.SetFlags(0)

	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("need 1 arg: timestamp file")
	}

	f, errOpen := os.Open(flag.Arg(0))
	Check(errOpen)
	defer f.Close()

	hist := make([]float64, *flag_n)

	t, g, err := ScanLine(f)
	last := 0.0
	for err != io.EOF {
		Check(err)

		if g == *flag_player {
			last = t
		} else {
			deltaT := t - last
			i := int(deltaT / (*flag_res))
			if i < len(hist) {
				hist[i] += 1 / (*flag_res)
			}
		}

		t, g, err = ScanLine(f)
	}

	time := make([]float64, *flag_n)
	for i := range time {
		time[i] = *flag_res * float64(i)
	}

	PrintTable(time, hist)
}
