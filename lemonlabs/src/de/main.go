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
		De{5},
		6,
	}
	fmt.Printf("%v , %T\n", d2, d2)
	n2 := d2.Lance()
	fmt.Printf("face :%v", n2)
	// 100 tirages
	var tirages1 [100]int
	for i := 0; i < 100; i++ {
		tirages1[i] = d1.Lance()
	}

}
