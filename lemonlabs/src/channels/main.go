package main

import (
	"fmt"
	"time"
)

type monchanel chan int

func pong(id int, channel chan int) {
	fmt.Printf("Dans %v\n", id)
	for j := range channel {
		fmt.Printf("id: %v  recu: %v\n", id, j)
		//time.Sleep(time.Duration(id) * time.Second)
	}
	fmt.Printf("id: %v traitement terminé\n", id)
}
func main() {
	//c := make(monchanel)
	for j := 1; j < 20; j++ {
		go func(id int) { fmt.Println(id) }(j)
	}
	/*
		for j := 1; j < 5; j++ {
			z := j
			go pong(z, c)
		}
		for i := 0; i < 100; i++ {
			c <- i
			//time.Sleep(10 * time.Millisecond)
		}
		close(c)
		time.Sleep(100 * time.Millisecond)
	*/
	time.Sleep(1000 * time.Millisecond)
}
