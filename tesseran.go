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


	t, g, err := scanLine(f)
	for err != io.EOF{
		check(err)
		fmt.Println(t, g)	
		t, g, err = scanLine(f)
	}
}

func scanLine(f io.Reader) (time float64, gotfragged bool, err error){
	var t float64
	var g bool
	_, err = fmt.Fscan(f, &t)
	if err != nil{
		return 0, false, err
	}
	_, err = fmt.Fscan(f, &gotfragged)
	if err != nil{
		return 0, false, err
	}
	return t, g, err
}

func check(err error){
		if err != nil{
			log.Fatal(err)
		}
}