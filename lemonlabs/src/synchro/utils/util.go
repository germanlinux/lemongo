package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LastPath(chaine string) string {
	//dir, file := filepath.Split("chaine")
	parts := strings.Split(chaine, string(os.PathSeparator))
	if len(parts) == 1 {
		//fmt.Println("EEE3", chaine)
		return chaine
	}
	fmt.Println("EEE4", parts)
	if parts[len(parts)-1]  == "" {
		//fmt.Println(("je passe "))
		return "." + string(os.PathSeparator)
	} 
	return parts[len(parts)-1]
}
func ParsePath(chaine, base string) (string, string, int) {
	dir, file := filepath.Split(chaine)
	prof := 0
	//fmt.Printf("on mac  %v, %v\n", dir, file)
	if dir == "" {
		prof  = 1
		dir = "." + string(os.PathSeparator)
		return dir, file, prof
	}
	//fmt.Printf("on mac2 %v, %v, %v\n", dir, file, base)
	parts := strings.Split(dir, base)
	//fmt.Println("part:", parts)
	pathf := ""
	if len(parts) > 1 {
		if parts[1] == "/" {
			pathf = base + "/"
		} else {
			pathf = base + parts[1]

		}
	} else {
		pathf = dir
	}
    prof = len(parts)
	return pathf, file, prof

}
