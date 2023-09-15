package main

type monchanel <-chan int
func main {
c := make(monchanel)
d := make(monchanel)
go client(1, c)
go client(2, d)
}