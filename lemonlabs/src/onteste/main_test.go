package main

import "testing"

func TestCalcul(t *testing.T) {

	want := 12
	msg := calcul(3, 4)
	if want != msg {
		t.Fatalf("calcul(3,4) = %v, want match for 13", msg)
	}
}
