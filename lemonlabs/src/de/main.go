package main

import "fmt"

type De = struct{ Nbface int }

func main() {

	d1 := De{6}
	fmt.Printf("%v , %T", d1, d1)
}
