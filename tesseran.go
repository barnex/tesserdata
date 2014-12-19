// Statiscal analysis of output generated by tesserstamp.
package main

import(
	"flag"
	"io"
	"os"
	"log"
	"fmt"
)

func main(){
	log.SetFlags(0)

	flag.Parse()
	if flag.NArg() != 1{
		log.Fatal("need 1 arg: timestamp file")
	}

	f, errOpen := os.Open(flag.Arg(0))
	check(errOpen)
	defer f.Close()

	a, b, _ := scanStamps(f)
	log.Println("you got", len(a), "frags, other player got", len(b))
}

// Reads timestamp file and return frag timestaps for you, other player and  total
func scanStamps(f io.Reader) (you, other, total []float64){
	t, g, err := scanLine(f)
	for err != io.EOF{
		check(err)
		total = append(total, t)
		if g{
			other = append(other, t)
		}else{
			you = append(you, t)
		}
		t, g, err = scanLine(f)
	}
	return
}

// Scan a single line from timestamp file.
func scanLine(f io.Reader) (time float64, gotfragged bool, err error){
	var t float64
	var g bool
	_, err = fmt.Fscan(f, &t)
	if err != nil{
		return 0, false, err
	}
	_, err = fmt.Fscan(f, &g)
	if err != nil{
		return 0, false, err
	}
	return t, g, err
}

// fatal if error != nil
func check(err error){
		if err != nil{
			log.Fatal(err)
		}
}
