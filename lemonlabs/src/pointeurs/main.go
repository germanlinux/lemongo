package main

import "fmt"

type Personne struct {
	Nom    string
	Prenom string
}

func New_personne(nom string) *Personne {
	return &Personne{Nom: nom}
}

func (p *Personne) Affiche() string {
	return p.Nom

}
func main() {
	s1 := "toto"
	p1 := &s1
	p2 := p1
	s2 := *p2
	fmt.Printf("s1 %v   p1 %v, *p1 %v s2 %v, p2 %v\n", s1, p1, *p1, s2, p2)
	s1 = s1 + "er"
	s2 = *p2
	fmt.Printf("s1 %v   p1 %v, *p1 %v s2 %v\n", s1, p1, *p1, s2)
	fmt.Println("--------------------------------")
	var per1 Personne
	per1.Nom = "Dupont"
	per1.Prenom = "jean"
	fmt.Printf("per1:%v \n", per1)
	var per2 Personne = Personne{Nom: "Dupond"}
	fmt.Printf("per2:%v \n", per2)
	var per3 Personne = Personne{"Dupond", "emile"}
	fmt.Printf("per3:%v \n", per3)
	fmt.Println(per2.Affiche())
	var per4 = New_personne("Robin")
	fmt.Printf("per4:%v \n", *per4)
	type S0 = struct {
		N1 int `essai`
		Y  bool
	}
	var monstr S0
	fmt.Printf("monstr:%v\n", monstr)
	fmt.Println("---------------------------------")
	mymap := make(map[string]int)
	mymap["un"] = 1
	mymap["deux"] = 2
	mymap["quatre"] = 4
	mymap["six"] = 6
	fmt.Println(mymap)
	for cle := range mymap {
		if cle == "deux" {
			delete(mymap, cle)
		}
	}
	// err := delete(mymap, "deux")
	// fmt.Println(mymap, err)
}
