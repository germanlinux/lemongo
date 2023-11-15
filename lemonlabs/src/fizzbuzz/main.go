package main

import (
	"fmt"
	"os"
	"strconv"
)
func verif3(n int) bool {
	res := n%3
	return res == 0
} 

func verif5(n int) bool {
	res := n%5
	return res == 0
} 

func main() {
	tchiffre := os.Args[1]
	fmt.Printf("input:%v, %T\n", tchiffre, tchiffre)	
	nchiffre,err:= strconv.Atoi(tchiffre)
	if err != nil {
		fmt.Println("un entier SVP!")
		os.Exit(1)
	}
	flag :=0
	fmt.Printf("input:%v, %T\n", nchiffre, nchiffre)
    if verif3(nchiffre) {
		fmt.Printf("Fizz")
		flag =1
	}
    if verif5(nchiffre) {
		if flag == 1 {
		fmt.Println("buzz")
		flag = 1
		} else {
			fmt.Println("Buzz")
		}
	}	else {
		if flag ==1 {
			fmt.Printf("\n")
		}
	}
	if flag == 0 {	
		fmt.Println(tchiffre)
	}
}