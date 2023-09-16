package main

import (
	"fmt"
)

type monchanel chan int

func pong(id int, channel chan int) {
	for {
		j := <-channel
		fmt.Printf("id: %v  recu: %v\n", id, j)
		//time.Sleep(time.Duration(id) * time.Second)
	}
}
func main() {
	c := make(monchanel)

	go pong(1, c)
	go pong(2, c)
	go pong(3, c)
	go pong(4, c)
	for i := 0; i < 100; i++ {
		c <- i
		//time.Sleep(1 * time.Second)
	}
}
