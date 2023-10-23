package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("bonjour suite")
	d, err := io.ReadAll(r)
	if err != nil {
		fmt.Println(("erreur"))
		return
	}
	fmt.Println("eric", string(d))

	tranche_a := make([]int, 3, 6)
	tranche_a[0] = 3
	tranche_a[1] = 2
	fmt.Println("tranche_a", tranche_a, "capacité:", cap(tranche_a), "taille:", len(tranche_a))
	tranche_b := tranche_a[1:3]
	fmt.Println("tranche_b", tranche_b, "capacité:", cap(tranche_b), "taille:", len(tranche_b))

	//tranche_a = append(tranche_a, 23)
	fmt.Println("tranche_a", tranche_a, "capacité:", cap(tranche_a), "taille:", len(tranche_a))
	fmt.Println("tranche_b", tranche_b, "capacité:", cap(tranche_b), "taille:", len(tranche_b))
	//	fmt.Printf("%p a\n", &tranche_a)
	//	fmt.Printf("%p b\n", &tranche_b)
	tranche_b = append(tranche_b, 18)
	fmt.Println("tranche_a", tranche_a, "capacité:", cap(tranche_a), "taille:", len(tranche_a))
	fmt.Println("tranche_b", tranche_b, "capacité:", cap(tranche_b), "taille:", len(tranche_b))
	//
	tranche_b = append(tranche_b, 19)
	tranche_b = append(tranche_b, 20)
	tranche_b = append(tranche_b, 21)
	tranche_b = append(tranche_b, 22)
	tranche_b = append(tranche_b, 23)
	tranche_b[0] = 99
	fmt.Println("tranche_a", tranche_a, "capacité:", cap(tranche_a), "taille:", len(tranche_a))
	fmt.Println("tranche_b", tranche_b, "capacité:", cap(tranche_b), "taille:", len(tranche_b))

	lettres := []string{"etoile", "des", "neiges"}
	fmt.Printf("Adresse de la variable lettre: %p\n", &lettres)
	fmt.Println("lettres", lettres, "capacité:", cap(lettres), "taille:", len(lettres))
	sub1 := lettres[:2]
	sub2 := sub1[:]
	fmt.Println("sub1", sub1, "capacité:", cap(sub1), "taille:", len(sub1))
	fmt.Printf("%p sub1\n", &sub1)
	fmt.Printf("Adresse de la variable sub1: %p\n", &sub1)
	fmt.Println("sub2", sub2, "capacité:", cap(sub2), "taille:", len(sub2), &sub2)
	fmt.Printf("Adresse de la variable sub1: %p\n", &sub2)
	sub2[1] = "toto"
	fmt.Println("lettres", lettres, "capacité:", cap(lettres), "taille:", len(lettres))
	fmt.Println("sub1", sub1, "capacité:", cap(sub1), "taille:", len(sub1))
	fmt.Println("sub2", sub2, "capacité:", cap(sub2), "taille:", len(sub2))
	lettres = append(lettres, "pourquoi")
	lettres = append(lettres, "encore")

	fmt.Printf("Adresse de la variable lettre: %p\n", &lettres)
	fmt.Println("lettres", lettres, "capacité:", cap(lettres), "taille:", len(lettres))
	sub3 := lettres[0:2]
	fmt.Println("sub3", sub3, "capacité:", cap(sub3), "taille:", len(sub3))
	sub4 := sub3[:]
	sub4 = append(sub4, "toujours")
	fmt.Println("sub4", sub4, "capacité:", cap(sub4), "taille:", len(sub4))
	sub4[2] = "plus_toujours"
	fmt.Println("sub4", sub4, "capacité:", cap(sub4), "taille:", len(sub4))
	sub4[0] = "python"
	fmt.Println("sub4", sub4, "capacité:", cap(sub4), "taille:", len(sub4))
	fmt.Println("sub3", sub3, "capacité:", cap(sub3), "taille:", len(sub3))
	fmt.Println("lettres", lettres, "capacité:", cap(lettres), "taille:", len(lettres))
	tranche_c := make([]int, 3, 6)
	tranche_d := tranche_c[1:3:3]
	tranche_d[1] = 15
	tranche_d[0] = 99
	fmt.Println("tranche_c", tranche_c, "capacité:", cap(tranche_c), "taille:", len(tranche_c))
	fmt.Println("tranche_d", tranche_d, "capacité:", cap(tranche_d), "taille:", len(tranche_d))
}
