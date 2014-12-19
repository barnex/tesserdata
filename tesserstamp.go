//+build ignore

// Records frag time stamps form tesseract stdout. Usage:
// 	tesseract | tesserstamp > timestamps.txt
// Output: time stamp in ms, boolean indication if you got fragged
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	start := time.Now()
	for {
		l, _, err := in.ReadLine()
		if err != nil {
			break
		}
		line := string(l)
		if strings.Contains(line, "fragged") {
			gotfragged := strings.Contains(line, "got")
			stamp := time.Since(start).Seconds() * 1000
			fmt.Println(stamp, gotfragged)
		}
	}
}
