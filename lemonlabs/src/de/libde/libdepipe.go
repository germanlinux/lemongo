package libde

import (
	"math/rand"
)

type DePipe struct {
	Nbface   int
	Faceplus int
}

func (d DePipe) GetFace() int {
	return d.Nbface
}
func (d DePipe) Lance() int {
	n := rand.Intn(d.Nbface-1) + 1
	return n
}
