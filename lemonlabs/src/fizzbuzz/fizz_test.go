package main

import (
	"testing"
)

func BenchmarkFizz10(b *testing.B) {
	// run the Fizz function b.N times
	for n := 0; n < b.N; n++ {
			Fizzbuzz("35")
	}
}