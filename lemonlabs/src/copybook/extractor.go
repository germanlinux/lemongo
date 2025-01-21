package main

import (
	"regexp"
	"strconv"
	"strings"
)

// exxtractLevel : extraction du niveau d'une ligne de zonage
func extractLevel(ligne string) int {
	// suppress space
	  ligne= strings.TrimLeft(ligne, " ")
	  renum  := regexp.MustCompile(`^\d+`)
	  tup:= renum.FindStringIndex(ligne)
	  snumligne := ligne[tup[0]:tup[1] ]
	  numligne, _ := strconv.Atoi(snumligne)
	 return numligne
}

// extractPic : extraction de la clause picture du zonage
func extractPic(ligne string) string {
	tligne := util_reduce_space(ligne)
	format := ""
	for cp, cont := range tligne  {
		ind:= strings.Index(cont, "PIC") 
		if ind > -1 {	
			format = tligne[cp +1]
			lg := len(format) 
			if format[lg -1] == '.' {
				format =  format[0:lg-1]
			}		
		}
	}
	return format
}
// extractLgExt : extraction de la longueur externe d'une donnée
func extractLgExt(pic string) int {
	var nb int
	if pic[0] == 'X' {
		//compte les X 
		nb = util_count(pic, "X")
	}	
	if 	pic[0] == '9' {
		nb = util_count(pic, "9")
	}
	return nb
}
// extractName : extraction du nom de la variable
func extractName(ligne string) string {
	tligne := util_reduce_space(ligne)
	name := tligne[1]
	if name[len(name)-1] == '.' {
		name = name[0:len(name)-1]
	} 
	return name
}

// extractType : extraction du type de donnée
func extractType(pic string) string {
	typeof := "ind"
	if pic ==""  {
		pic = " "
	}
	tp := pic[0]
	switch tp {
	case 'X':
	  typeof = "ALN"
	case '9':
	  typeof = "NUM"	
	case 'S':
	  typeof = "SNUM"
	default: 
	  typeof = "GRP"  		
	}
	return typeof
}