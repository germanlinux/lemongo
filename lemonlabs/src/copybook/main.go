package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Variable struct {
	uuid       uuid.UUID
	name       string
	level      int
	typeof     string
	lg_externe int
	lg_entier  int
	lg_decimal int
	picture    string
	parentUuid uuid.UUID
	rank       int
	line       string
	sibling    uuid.UUID
}

// type Copybook []Variable
type Parent struct {
	niveau int
	uuid   uuid.UUID
}

func read_copybook(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		error := err
		return nil, error
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

/*
func EncodeBase64(chaine string) {

	input := []byte(chaine)
	myString := string(input[:])
	fmt.Println(myString)
	encoder := base64.StdEncoding.EncodeToString(input)
	//encoder.Write(input)
	//encoder.Close()
	fmt.Println("encode", encoder)
}
*/

func parseData(tlignes []string) []Variable {
	var table_niveau []Parent
	var result []Variable
	var typeof string
	nivparent := Parent{niveau: 0}
	table_niveau = append(table_niveau, nivparent)

	for cp, ligne := range tlignes {
		uuid := uuid.New()
		nom := extractName(ligne)
		niveau := extractLevel(ligne)
		picture := extractPic(ligne)
		if niveau == 88 {
			typeof = "COND"
		} else {
			typeof = extractType(picture)
		}
		table_niveau = append(table_niveau, Parent{niveau: niveau, uuid: uuid})
		uuidparent := verif_niveau(table_niveau)
		tmpv := Variable{uuid: uuid, name: nom, level: niveau, typeof: typeof, picture: picture, line: ligne, rank: cp + 1, parentUuid: uuidparent}
		if typeof == "GRP" || typeof == "COND" {
			tmpv.lg_externe = 0

		} else {
			tmpv.lg_externe = extractLgExt(picture)

		}
		if isFloat(picture) {
			lg1, lg2 := ln4float(picture)
			tmpv.lg_externe = lg1 + lg2
			tmpv.lg_entier = lg1
			tmpv.lg_decimal = lg2
		}
		result = append(result, tmpv)

		//EncodeBase64("encode moi vite")
		//	    fmt.Println(table_niveau)

		//	lg_intere := extractLgInt(picture)
		//	rank := cp +1
	}
	return result
}

func recherche_parent(pere uuid.UUID, table *[]Variable) (int, error) {
	for cp, item := range *table {
		if item.uuid == pere {
			return cp, nil
		}

	}

	return 0, errors.New("entite racine")
}

func main() {
	//uuid.New()
	filename := "data1.cob"
	lignes, err := read_copybook(filename)
	if err != nil {
		panic(err)
	}
	data := parseData(lignes)
	for _, item := range data {
		fmt.Println(item.name, item.typeof, item.lg_externe, item.uuid, item.level, item.parentUuid)

	}
	calcul_lg(&data)
	fmt.Println("------------")
	for cpx, item := range data {
		fmt.Printf("%v;%v;%v;%v\n", cpx, item.level, item.name, item.lg_externe)
	}
	fmt.Println("------------")
	a, b := init_grp(data)

	fmt.Println(b)
	fmt.Println("----------")
	fmt.Println(a)
	display(a, b)
	voisin(&data)

}
