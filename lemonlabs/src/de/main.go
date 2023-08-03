package main

import (
	"fmt"
	"math/rand"
)

type De struct{ Nbface int }

func (d De) get_face() int {
	return d.Nbface
}
func (d De) lance() int {
	n := rand.Intn(5) + 1
	return n
}

func main() {

	d1 := De{6}
	fmt.Printf("%v , %T\n", d1, d1)
	fmt.Printf("nb de face :%v\n", d1.get_face())
	n1 := d1.lance()
	fmt.Printf("face :%v", n1)
}
