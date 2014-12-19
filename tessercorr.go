//+build ignore

// Auto- and cross-correlation of frag times
package main

import (
	"flag"
	"io"
	"log"
	"os"

	. "."
)

var (
	flag_res = flag.Float64("res", 1, "Time resolution in seconds")
	flag_len = flag.Float64("len", 30, "Output duration in seconds")
	flag_a   = flag.Bool("a", false, "first input (true=you got fragged)")
	flag_b   = flag.Bool("b", false, "second input (true=you got fragged)")
)

func main() {
	log.SetFlags(0)

	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("need 1 arg: timestamp file")
	}
	res := *flag_res

	f, errOpen := os.Open(flag.Arg(0))
	Check(errOpen)
	defer f.Close()

	var time []float64
	var fragged []bool
	t, g, err := ScanLine(f)
	for err != io.EOF {
		Check(err)
		time = append(time, t)
		fragged = append(fragged, g)
		t, g, err = ScanLine(f)
	}

	corr := make([]float64, int(*flag_len/res))

	for i := range time {
		if fragged[i] != *flag_a {
			continue
		}
		for j := i; j < len(time); j++ {
			if fragged[j] != *flag_b {
				continue
			}
			deltaT := time[j] - time[i]
			if deltaT > *flag_len {
				break
			}
			i := int(deltaT / res)
			corr[i] += 1 / res
		}
	}

	dt := make([]float64, len(corr))
	for i := range dt {
		dt[i] = float64(i) * (res)
	}

	PrintTable(dt, corr)
}
