package libde

import (
	"math/rand"
)

type De struct{ Nbface int }

type DePipe struct {
	De
	Faceplus int
}

func (d De) GetFace() int {
	return d.Nbface
}
func (d De) Lance() int {
	n := rand.Intn(d.Nbface) + 1
	return n
}

func (d DePipe) GetFace() int {
	return d.De.Nbface
}
func (d DePipe) Lance() int {
	n := rand.Intn(d.De.Nbface+1) + 1
	if n == d.De.Nbface+1 {
		//fmt.Println("je passe")
		n = d.Faceplus
	}
	return n
}
