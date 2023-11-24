package util

import (
	"fmt"
	"path/filepath"
	"strings"
)

func ParsePath(chaine, base  string )(string, string) {
	dir, file := filepath.Split("chaine")
	fmt.Printf("on mac  %v, %v\n",dir, file)
    parts := strings.Split(dir, base)
    fmt.Println((parts))
	pathf:=""
	if len(parts) > 0 {
		if parts[1] == "/"  {
			pathf = base  + "/"		
		} else {
			pathf = base + parts[1]

		}
	} else {
		pathf = dir	
	}
	
	return pathf, file

}