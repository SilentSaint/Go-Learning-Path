package main

//echo command clone - prints it command line args
import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing new lines")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse() // must call before using flags
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
