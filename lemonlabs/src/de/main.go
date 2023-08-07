package main

import (
	"fmt"
	"lemonlabs/src/de/libde"
)

type De = libde.De
type DePipe = libde.DePipe

func main() {

	d1 := De{6}
	fmt.Printf("%v , %T\n", d1, d1)
	fmt.Printf("nb de face :%v\n", d1.GetFace())
	n1 := d1.Lance()
	fmt.Printf("face :%v", n1)
	d2 := DePipe{
		De{6},
		6,
	}
	fmt.Printf("%v , %T\n", d2, d2)
	n2 := d2.Lance()
	fmt.Printf("face :%v\n", n2)
	// 100 tirages
	tirages1 := make(map[int]int)
	for i := 0; i < 100; i++ {
		z := d1.Lance()
		tirages1[z] += 1
	}
	for face, score := range tirages1 {
		fmt.Println("face:", face, "score:", score)
	}
	fmt.Printf("face plus: %v\n", d2.Faceplus)
	tirages2 := make(map[int]int)
	for i := 0; i < 1000; i++ {
		z := d2.Lance()
		tirages2[z] += 1
	}
	fmt.Println("-------------------------------------")
	for face, score := range tirages2 {
		fmt.Println("face:", face, "score:", score)
	}
}
