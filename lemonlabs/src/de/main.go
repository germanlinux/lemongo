package main

import (
	"fmt"
	"lemonlabs/src/de/libde"
	"sort"
)

type De = libde.De
type DePipe = libde.DePipe
type DePipe2 = libde.DePipe2

type Maliste []any

func (a Maliste) Len() int      { return len(a) }
func (a Maliste) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Maliste) Less(i, j int) bool {
	si := fmt.Sprintf("%v", a[i])
	sj := fmt.Sprintf("%v", a[j])
	if _, ok := a[i].(int); ok {
		si = fmt.Sprintf("|%v", a[i])
	}
	if _, ok := a[j].(int); ok {
		sj = fmt.Sprintf("|%v", a[j])
	}

	return si < sj
}
func main() {

	d1 := De{6}
	fmt.Printf("%v , %T\n", d1, d1)
	fmt.Printf("nb de face (adresse):%v\n", d1.GetFace())
	fmt.Printf("nb de face (pointeur):%v\n", (&d1).GetFace())
	d1p := &d1
	fmt.Printf("nb de face (pointeur V2):%v\n", d1p.GetFace())

	n1 := d1.Lance()
	fmt.Printf("face :%v\n", n1)
	d2 := DePipe{
		De{6},
		6,
	}
	fmt.Printf("%v , %T\n", d2, d2)
	n2 := d2.Lance()
	fmt.Printf("face :%v\n", n2)
	// 100 tirages
	tirages1 := make(map[int]int)
	for i := 0; i < 100; i++ {
		z := d1.Lance()
		tirages1[z] += 1
	}
	for face, score := range tirages1 {
		fmt.Println("face:", face, "score:", score)
	}

	fmt.Println("-------------------------------------")
	fmt.Printf("face plus: %v\n", d2.Faceplus)
	tirages2 := make(map[int]int)
	for i := 0; i < 1000; i++ {
		z := d2.Lance()
		tirages2[z] += 1
	}
	fmt.Println("-------------------------------------")
	var pourcent float64

	for face, score := range tirages2 {
		pourcent = float64(score)
		pourcent = pourcent / 10
		s := fmt.Sprintf("%.2f", pourcent)
		fmt.Println("face:", face, "score:", s)
	}
	fmt.Println("Tri par face")
	tface := make([]int, 0)
	for key := range tirages2 {
		tface = append(tface, key)
	}
	fmt.Println(tface)
	sort.Ints(tface)
	fmt.Println(tface)
	fmt.Println("-----------tri sur la face --------------------------")
	for _, key := range tface {

		pourcent = float64(tirages2[key])
		pourcent = pourcent / 10
		s := fmt.Sprintf("%.2f", pourcent)
		fmt.Println("face:", key, "score:", s)
	}
	fmt.Println("--------tri par valeur ------")
	sort.SliceStable(tface, func(i, j int) bool { return tirages2[tface[i]] < tirages2[tface[j]] })
	fmt.Println("par valeur", tface)
	for _, key := range tface {
		pourcent = float64(tirages2[key])
		pourcent = pourcent / 10
		s := fmt.Sprintf("%.2f", pourcent)
		fmt.Println("face:", key, "score:", s)
	}
	fmt.Println("-----------tri sur la face ordre inverse --------------------------")
	sort.SliceStable(tface, func(i, j int) bool { return tirages2[tface[i]] >= tirages2[tface[j]] })
	fmt.Println("par valeur", tface)
	for _, key := range tface {
		pourcent = float64(tirages2[key])
		pourcent = pourcent / 10
		s := fmt.Sprintf("%.2f", pourcent)
		fmt.Println("face:", key, "score:", s)
	}
	d2p := &d2
	fmt.Println("pointeur", d2p.Lance())
	// appel de fonction par adresse ou par pointeur
	var d3 De
	d3 = De{12}
	fmt.Println("par valeur:", d3.GetFaceValeur())
	var dp *De
	dp = &d3
	fmt.Println("par pointeur:", dp.GetFacePointeur())
	// 2 formes possibles
	fmt.Println("recepteur pointeur:", d3.GetFacePointeur())
	fmt.Println("recepteur valeur:", dp.GetFaceValeur())
	// methode setter
	d4 := DePipe{
		De{6},
		5,
	}
	fmt.Println("avant appel", d4)
	d4.SetFaceValeur(2)
	fmt.Println("apres appel", d4)
	tmp := &d4
	tmp.SetFacePointeur(3)
	fmt.Println("apres appel", d4)
	(&d4).SetFacePointeur(6)
	fmt.Println("apres appel", d4)
	d4.SetFacePointeur(1)
	fmt.Println("apres appel", d4)
	d4.SetFaceDEPointeur(10)
	fmt.Println("apres appel", d4)
	fmt.Println("complet", d4.De.Nbface)
	fmt.Println("direct", d4.Nbface)
	fmt.Println("essai pointeur stryc imbriquée")
	d2spe := DePipe2{
		&De{6},
		6,
	}
	fmt.Printf("ok : %v\n", d2spe)
	fmt.Printf("face %v\n", d2spe.De.Nbface)

	a := []int{3, 6, 4, 1}
	fmt.Println("avant", a)
	sort.Ints(a)
	fmt.Println("après", a)
	g := Maliste{3, 6, "a", 4, 1}
	fmt.Printf("g avant tri: %T, %v\n", g, g)
	//	sort.SliceStable(g, func(i, j int) bool { )
	//	fmt.Printf("g après tri: %T, %v\n", g, g)
	sort.Sort(g)
	fmt.Printf("g apres tri: %T, %v\n", g, g)
	sort.Sort(sort.Reverse(g))
	fmt.Printf("g apres inversion tri: %T, %v\n", g, g)

}
