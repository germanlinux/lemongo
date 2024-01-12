package main

import (
	"fmt"
	"strings"
)

func main() {
	chaine := "14"
	s:= make([]string,10)
	if chaine != "" {
		s = strings.Split(chaine, ",")
		fmt.Println(s)

	}
	for i, id := range s {
		id := strings.Trim(id," ")
		if id != "" {
		fmt.Println(i, id)
		}
	}
}
