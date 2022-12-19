package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// this program prints the data found at the current url

func main() {
	for _, url := range os.Args[1:] {
		//Use http.Get method to fetch an URL in go
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
			os.Exit(1)
		}
		//since response is a readable stream   we are using ReadAll method
		// or
		// we could use io.copy to directly write to the output stream
		_, err = io.Copy(os.Stdout, res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%v", res.Status)

	}
}
