package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	//fmt.Println(line)
	if !strings.Contains(line, "Go") {
		return false, line, 0
	}
	compteur := strings.Count(line, "Go")
	nwligne := strings.ReplaceAll(line, "Go", "Python")
	return true, nwligne, compteur
}

func FindRemplaceFile(src, dst, old, new string) (occ int, line []int, err error) {
	fmt.Println("ligne entree", line)
	compteur := 0
	file, err := os.Open(src)
	defer file.Close()
	dfile, err2 := os.Create(dst)
	defer dfile.Close()
	if err2 != nil {
		return occ, line, err2
	}
	//handle errors while opening
	if err != nil {
		return occ, line, err
	}

	/*
		scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}*/

	fileScanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(dfile)
	defer writer.Flush()

	// read line by line
	cpligne := 0
	for fileScanner.Scan() {
		cpligne += 1
		rep, chaine, occ := ProcessLine(fileScanner.Text(), old, new)
		if rep {
			compteur += occ
			line = append(line, cpligne)
		}
		chaine += "\n"
		fmt.Fprint(writer, chaine)
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		fmt.Printf("Error while reading file: %s", err)
	}
	return compteur, line, err
}
func main() {
	nbocc, numligne, erreur := FindRemplaceFile("wikigo.txt", "python.txt", "Go", "Python")
	if erreur != nil {
		fmt.Println("erreur generale")
	}
	fmt.Println(nbocc, numligne)
}
