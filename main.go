package main

import (
	// used for easier input and output
	"bufio"
	"fmt"
	"os"
)

/*
In v2 of this program the bufio works in streaming mode ie. the input is read line by line
*/
func main() {
	//make creates a new empty map
	files := os.Args[1:]
	counts := make(map[string]int)
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
