package main

import (
	"fmt"
	"time"
)

type monchanel chan int

func pong(id int, channel chan int) {
	for j := range channel {
		fmt.Printf("id: %v  recu: %v\n", id, j)
		//time.Sleep(time.Duration(id) * time.Second)
	}
	fmt.Printf("id: %v traitement terminé\n", id)
}
func main() {
	c := make(monchanel)

	go pong(1, c)
	go pong(2, c)
	go pong(3, c)
	go pong(4, c)
	for i := 0; i < 100; i++ {
		c <- i
		//time.Sleep(10 * time.Millisecond)
	}
	close(c)
	time.Sleep(100 * time.Millisecond)
}
