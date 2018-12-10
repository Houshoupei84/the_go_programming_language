package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false , "omit trailing newline")
var sep = flag.String("s", " ", "separator")

/*
testcase

./echo4 a bc def


*/
func main(){

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n{
		fmt.Println()
	}
}