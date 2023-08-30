package main

import (
	"io/ioutil"
	"log"
)

func readfilehtml(filename string) (string, error) {
	bst, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
}

func main() {
	filename := "page.html"
	text, err := readfilehtml(filename)
	if err != nil {
		log.Fatal(err)
	}

}
