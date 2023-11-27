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
		return chaine
	}
	fmt.Println("parts last", parts)
	return parts[len(parts)-1]
}
func ParsePath(chaine, base string) (string, string) {
	dir, file := filepath.Split(chaine)
	fmt.Printf("on mac  %v, %v\n", dir, file)
	if dir == "" {
		dir = "." + string(os.PathSeparator)
		return dir, file
	}
	fmt.Printf("on mac  %v, %v\n", dir, file)
	parts := strings.Split(dir, base)
	fmt.Println("part:", parts)
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

	return pathf, file

}
