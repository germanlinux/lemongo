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

func verif5(tn string) bool {
	if tn[len(tn) - 1] == '0' ||  tn[len(tn) - 1]== '5' {
	return true
	} 
	return false
}
func Fizzbuzz(tchiffre string) string {
	var chaine string
	nchiffre,err:= strconv.Atoi(tchiffre)
	if err != nil {
		fmt.Println("un entier SVP!")
		os.Exit(1)
	}
	flag :=0
	//fmt.Printf("input:%v, %T\n", nchiffre, nchiffre)
    if verif3(nchiffre) {
		chaine ="Fizz"
		flag =1
	}
    if verif5(tchiffre) {
		if flag == 1 {
		chaine += "buzz"
		flag = 1
		} else {
			chaine ="Buzz"
			flag =1	
		}
	}	
	if flag == 0 {	
		chaine = tchiffre
	}
	return chaine

}
func main() {
	tchiffre := os.Args[1]
	//fmt.Printf("input:%v, %T\n", tchiffre, tchiffre)	
	fmt.Println(Fizzbuzz(tchiffre))
}