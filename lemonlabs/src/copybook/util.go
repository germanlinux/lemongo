package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// this function split format on the "V" in the picture
func util_split_float(pic string) (int, int) {
	tup := strings.Split(pic, "V")
	deb := tup[0]
	fin := tup[1]
	lg1 := extractLgExt(deb)
	lg2 := extractLgExt(fin)
	return lg1, lg2
}

// this function split line on space and delete item of empty string
func util_reduce_space(ligne string) []string {
	tligne := strings.Split(ligne, " ")
	var tlignenw []string
	for _, cont := range tligne {
		if cont != "" {
			tlignenw = append(tlignenw, cont)
		}
	}
	return tlignenw
}
func util_count(pic string, pattern string) int {
	re_parant := regexp.MustCompile(`(\(\d+\))`)
	var nb int

	if strings.Count(pic, "(") > 0 {
		tup := re_parant.FindSubmatch([]byte(pic))
		if len(tup) > 1 {
			chaine := string(tup[0])
			chaine = chaine[1 : len(chaine)-1]
			tnb, _ := strconv.Atoi(chaine)
			nb = tnb
		}
	} else {
		nb = strings.Count(pic, pattern)

	}
	return nb
}

// isFloat :  retourne vrai ou faux en fonction de la presence d'une déclaration de virgule flottante
func isFloat(pic string) bool {
	return strings.Contains(pic, "V")
}

// isSigned :  retourne vrai ou faux en fonction de la presence d'une déclaration de type S9..
func isSigned(pic string) bool {
	return strings.Contains(pic, "S9")
}

// ln4float :  calcule la longueur fixe et decimale
func ln4float(pic string) (fixe int, dec int) {
	if isSigned(pic) {
		pic = pic[1:]
	}
	fixe, dec = util_split_float(pic)
	//	tmpv.lg_externe = lg1 + lg2
	return fixe, dec
}

// verif_niveau : retourne uuid de la variable parent
func verif_niveau(table []Parent) uuid.UUID {
	lg := len(table)
	dernier := table[lg-1]
	if lg-1 == 0 {
		return dernier.uuid
	}

	for cp := lg; cp > 0; cp-- {
		prec := table[cp-1]
		// cas simple
		if dernier.niveau > prec.niveau {
			return prec.uuid
		}
	}
	return uuid.Nil
}

// findAllson : retourne la liste des variables filles d'une variable de type groupe
func findAllson(item uuid.UUID, table []Variable) []uuid.UUID {
	var result []uuid.UUID
	for _, data := range table {
		if data.parentUuid == item {
			result = append(result, data.parentUuid)
		}
	}
	return result
}

// init_grp : construit l'arbre hierarchique des données groupe et d'un dictionnaire uuid/nom
func init_grp(table []Variable) (map[uuid.UUID][]uuid.UUID, map[uuid.UUID]string) {

	dictA := make(map[uuid.UUID]string, 0)
	dictGRP := make(map[uuid.UUID][]uuid.UUID, 0)
	for _, item := range table {
		dictA[item.uuid] = item.name
	}
	for _, item := range table {
		if item.typeof == "GRP" {
			listefils := findAllson(item.uuid, table)
			dictGRP[item.uuid] = listefils
		}
	}
	return dictGRP, dictA
}

// func calcul_lg : calcule les longueurs externe des zones groupes
// on parcours une la table de variable triée par niveau (decroissant)  pour cumuler les zones simples et on remonte dans la hierarchie.
func calcul_lg(table *[]Variable) {
	//var change []Variable

	//trier par niveau
	tmp_change := *table
	sort.Slice(tmp_change, func(i, j int) bool {
		return tmp_change[i].level > tmp_change[j].level
	})

	for _, item := range tmp_change {
		index, err := recherche_parent(item.parentUuid, table)
		if err == nil {
			(*table)[index].lg_externe += item.lg_externe
		}
	}
}

// ajouter le voisin
func voisin(table *[]Variable) {
	//var change []Variable
	debut := 0
	fin := 0
	lg := 0
	for -,item := range table {
		if item.typeof == "GRP" || item.typeof == "COND" {
			lg = 0
		} else {
			lg = item.lg_externe
		}	
	}

}
func display(principal map[uuid.UUID][]uuid.UUID, dict map[uuid.UUID]string) {
	for cle, val := range principal {
		fmt.Println("cle", cle, dict[cle])
		fmt.Println(val)
	}

}
