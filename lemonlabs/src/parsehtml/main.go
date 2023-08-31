package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func readfilehtml(filename string) (string, error) {
	bst, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bst), err
}

/*
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
				fmt.Printf("token: %v", t)
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
*/
func main() {
	dico := make(map[string]int)
	filename := "page.html"
	text, err := readfilehtml(filename)
	if err != nil {
		log.Fatal(err)
	}
	//data := parse(text)
	//fmt.Println(data)
	pattern := regexp.MustCompile("[\"2]{1}&gt;</span>([^<]+)")
	allst := pattern.FindAllStringSubmatch(text, -1)
	for _, match := range allst {
		if len(match[1]) > 1 {
			//	fmt.Printf("%v\n", match[1])
			dico[match[1]] = 1
		}
	}
	for cle, _ := range dico {
		fmt.Println(cle)
	}
}
