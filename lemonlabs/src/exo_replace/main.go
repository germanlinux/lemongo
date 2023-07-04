package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	if !strings.Contains(line, "Go") {
		return false, "", 0
	}
	compteur := strings.Count(line, "Go")
	nwligne := strings.ReplaceAll(line, "Go", "Python")
	return true, nwligne, compteur
}

func FindRemplaceFile(src, old, new string) (occ int, line []int, err error) {
	fmt.Println("ligne entree", line)
	compteur := 0
	file, err := os.Open(src)
	defer file.Close()
	//handle errors while opening
	if err != nil {
		fmt.Printf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	// read line by line
	cpligne := 0
	for fileScanner.Scan() {
		cpligne += 1
		rep, chaine, occ := ProcessLine(fileScanner.Text(), old, new)
		if rep {
			compteur += occ
			line = append(line, cpligne)
			fmt.Println(chaine)

		}
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		fmt.Printf("Error while reading file: %s", err)
	}
	return compteur, line, err
}
func main() {
	nbocc, numligne, erreur := FindRemplaceFile("wikigo.txt", "Go", "Python")
	if erreur != nil {
		fmt.Println("erreur generale")
	}
	fmt.Println(nbocc, numligne)
}
