package main

import "fmt"

func itoa(n int) string {
	return fmt.Sprintf("%v = %d =  %T", n, n, n)
}

type T int

func (t T) f() string {
	fmt.Printf(" x%Tx, %v", t, t)
	return itoa(int(t))

}
func main() {

	t := T(3)
	fmt.Println(itoa(2))
	fmt.Println(t.f())
	fmt.Println(T.f(4))
}
