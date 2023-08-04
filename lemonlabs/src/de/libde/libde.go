package libde

import (
	"math/rand"
)

type De struct{ Nbface int }

func (d De) GetFace() int {
	return d.Nbface
}
func (d De) Lance() int {
	n := rand.Intn(d.Nbface-1) + 1
	return n
}
