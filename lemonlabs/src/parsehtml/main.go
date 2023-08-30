package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func readfilehtml(filename string) (string, error) {
	bst, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bst), err
}

func parse(text string) (data []string) {
	tkn := html.NewTokenizer(strings.NewReader(text))
	var vals []string
	var isH3 bool
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:
			t := tkn.Token()
			isH3 = t.Data == "td"
		case tt == html.TextToken:
			t := tkn.Token()
			if isH3 {
				vals = append(vals, t.Data)
			}
			isH3 = false
		}
	}
}
func main() {
	filename := "page.html"
	text, err := readfilehtml(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := parse(text)
	fmt.Println(data)
}
