package main

import (
	// used for easier input and output

	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
In v2 of this program the bufio works in streaming mode ie. the input is read line by line
*/
func main() {
	//make creates a new empty map
	files := os.Args[1:]
	counts := make(map[string]int)
	for _, arg := range files {
		/* Readfile reads the entire file into memory, and returns a byte slice */
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// func countLines(f *os.File, counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// }
