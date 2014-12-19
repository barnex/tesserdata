// common functions for reading tesserstamp files
package tesserdata

import (
	"fmt"
	"io"
	"log"
)

// Prints slices as columns
func PrintTable(columns ...[]float64) {
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

// Make time column for stamps data and time resolution.
// Returned times go from 0 to the last time stamp in steps of res seconds.
func MakeTime(s []float64, res float64) []float64 {
	tmax := s[len(s)-1]
	imax := int(tmax / res)
	time := make([]float64, imax+1)
	for i := range time {
		time[i] = float64(i) * res
	}
	return time
}

// Reads timestamp file and return frag timestaps for you, other player and  total
func ScanStamps(f io.Reader) (you, other, total []float64) {
	t, g, err := ScanLine(f)
	for err != io.EOF {
		Check(err)
		total = append(total, t)
		if g {
			other = append(other, t)
		} else {
			you = append(you, t)
		}
		t, g, err = ScanLine(f)
	}
	return
}

// Scan a single line from timestamp file.
func ScanLine(f io.Reader) (time float64, gotfragged bool, err error) {
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
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
