// +build ignore

// Histogram of expected lifetimes
package main

import (
	"flag"
	"log"
	"os"

	. "."
)

var (
	flag_res = flag.Float64("res", 1, "Time resolution in seconds")
	flag_n   = flag.Int("bins", 100, "N bins")
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

	a, b, _ := ScanStamps(f)

	rateA := hist(a)
	rateB := hist(b)

	time := make([]float64, *flag_n)
	for i := range time {
		time[i] = *flag_res * float64(i)
	}

	PrintTable(time, rateA, rateB)
}

func hist(s []float64) []float64 {
	hist := make([]float64, *flag_n)

	for i := 1; i < len(s); i++ {
		di := int((s[i] - s[i-1]) / (*flag_res))
		if di < len(hist) {
			hist[di] += 1 / (*flag_res)
		}
	}
	return hist
}
