package main

import (
	"fmt"
	"lemonlabs/src/de/libde"
)

func main() {

	d1 := libde.De{6}
	fmt.Printf("%v , %T\n", d1, d1)
	fmt.Printf("nb de face :%v\n", d1.GetFace())
	n1 := d1.Lance()
	fmt.Printf("face :%v", n1)
	d2 := libde.DePipe{
		Nbface:   6,
		Faceplus: 6,
	}
	fmt.Printf("%v , %T\n", d2, d2)

}
