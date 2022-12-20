package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// wait demo's the retry - error handling strategy
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wait url\n")
		os.Exit(1)
	}
	url := os.Args[1]
	log.SetPrefix("wait: ")
	log.SetFlags(0)
	if err := waitForServer(url); err != nil {
		log.Fatalf("Site is down %v", err)
		os.Exit(1)
	}
}

func waitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying..", err)
		time.Sleep(time.Second << uint(tries)) //exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
