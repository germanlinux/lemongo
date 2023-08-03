package main

import "fmt"

type De struct{ Nbface int }

func (d De) get_face() int {
	return d.Nbface
}

func main() {

	d1 := De{6}
	fmt.Printf("%v , %T\n", d1, d1)
	fmt.Printf("nb de face :%v", d1.get_face())
}
