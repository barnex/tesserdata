//+build ignore

// Extract frag rates from time stamp file
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	flag_res = flag.Float64("res", 10, "Time resolution in seconds")
)

func main() {
	log.SetFlags(0)

	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("need 1 arg: timestamp file")
	}

	f, errOpen := os.Open(flag.Arg(0))
	check(errOpen)
	defer f.Close()

	a, b, total := scanStamps(f)

	rateA := fragRate(a, *flag_res)
	rateB := fragRate(b, *flag_res)
	time := makeTime(total, *flag_res)

	fmt.Println("# time(s) fragrate_A(Hz) fragrate_B(Hz)")
	printTable(time, rateA, rateB)
}

func printTable(columns ...[]float64) {
	t := columns[0]
	for i := range t {
		for c := range columns {
			v := 0.0
			if len(columns[c]) > i {
				v = columns[c][i]
			}
			fmt.Print(v, "\t")
		}
		fmt.Println()
	}
}

func makeTime(s []float64, res float64) []float64 {
	tmax := s[len(s)-1]
	imax := int(tmax / res)
	time := make([]float64, imax+1)
	for i := range time {
		time[i] = float64(i) * res
	}
	return time
}

func fragRate(s []float64, res float64) []float64 {
	tmax := s[len(s)-1]
	imax := int(tmax / res)
	rate := make([]float64, imax+1)
	weight := 1 / res
	for _, t := range s {
		i := int(t / res)
		rate[i] += weight
	}
	return rate
}

// Reads timestamp file and return frag timestaps for you, other player and  total
func scanStamps(f io.Reader) (you, other, total []float64) {
	t, g, err := scanLine(f)
	for err != io.EOF {
		check(err)
		total = append(total, t)
		if g {
			other = append(other, t)
		} else {
			you = append(you, t)
		}
		t, g, err = scanLine(f)
	}
	return
}

// Scan a single line from timestamp file.
func scanLine(f io.Reader) (time float64, gotfragged bool, err error) {
	var t float64
	var g bool
	_, err = fmt.Fscan(f, &t)
	if err != nil {
		return 0, false, err
	}
	_, err = fmt.Fscan(f, &g)
	if err != nil {
		return 0, false, err
	}
	return t / 1000, g, err
}

// fatal if error != nil
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}