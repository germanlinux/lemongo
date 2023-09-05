package libde

import (
	"fmt"
	"math/rand"
)

type De struct{ Nbface int }

type DePipe struct {
	De
	Faceplus int
}
type DePipe2 struct {
	*De
	Faceplus int
}

// GetFace: cette fonction/methode retourne le nombre de face du dé.
func (d De) GetFace() int {
	return d.Nbface
}

// Lance: cette fonction/méthode simule un lancer de dé et retourne le numero de la face (commence par 1)
func (d De) Lance() int {
	n := rand.Intn(d.Nbface) + 1
	return n
}

func (d De) GetFaceValeur() int {
	return d.Nbface
}
func (d *De) GetFacePointeur() int {
	return d.Nbface
}

func (d DePipe) GetFace() int {
	return d.De.Nbface
}

func (d DePipe) SetFaceValeur(face int) DePipe {
	d.Faceplus = face
	fmt.Println("par valeur", d)
	return d
}
func (d *DePipe) SetFacePointeur(face int) *DePipe {

	d.Faceplus = face
	fmt.Println("par pointeur", *d)
	return d
}
func (d *DePipe) SetFaceDEPointeur(face int) *DePipe {

	d.Nbface = face
	fmt.Println("par pointeur", *d)
	return d
}

func (d DePipe) Lance() int {
	n := rand.Intn(d.De.Nbface+1) + 1
	if n == d.De.Nbface+1 {
		n = d.Faceplus
	}
	return n
}
