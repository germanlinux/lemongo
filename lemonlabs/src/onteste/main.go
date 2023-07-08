package main

import "fmt"

func calcul(a, b int) (res int) {
	res = a*b + 1
	return res
}

func main() {
	res := calcul(3, 4)
	fmt.Println(res)
}
