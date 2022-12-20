package main

//this package aims to demonstrate a web crawler

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findLinks: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Print(link)
		}
	}

}

func visit(links []string, n *html.Node) []string {
	//check node type and if anchor node then go through attr's and find href tag
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	//recursion to traverse the whole tree
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func findLinks(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("getting %s : %s", url, res.Status)
	}
	doc, err := html.Parse(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as html: %v", url, err)
	}
	return visit(nil, doc), nil
}

//bare return example
// func countWordsAndImages(url string)(words,images int, err error) {
// 	resp,err := http.Get(url)
// 	if err != nil {
// 		return
// 	}
// 	doc,err != http.parse(resp.Body)
// 	resp.Body.Close()
// 	if err != nil {
// 		err = fmt.Errorf("parsing html : %s",err)
// 		return
// 	}
// 	words,images = countWordsAndImages(doc)
// 	return
// }

// func countWordsAndImages(n *html.Node)	(words,images int) {
// 	if n.Type == html.ElementNode {
// 		if n.Data == "img"{
// 			images++
// 		}
// 	}
// 	return
// }
