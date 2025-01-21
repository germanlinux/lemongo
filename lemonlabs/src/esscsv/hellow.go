package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/* var Inp = `"Sample Entry #2","Michael321","12345","https://keepass.info/help/kb/testform.html",""
"pgadmin","german","german","",""
"microsoft","germanliXXX.fr","XXXXXXXXX","",""
"zvm","root","rootvmeg","XXXXXXXXX",""
"collabra","XXXXXX.fr","XXXXXX","https://collabora.dgfip.finances.rie.gouv.fr/",""
`
*/
//"Account","Login Name","Password","Web Site","Comments"

type Entry struct {
	Name     string
	User     string
	Password string
	Etat     int
}
type Entries struct {
	File     string
	Mentries map[string]*Entry
}

func lire_fichier(nom string) Entries {
	filename := nom
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()
	filescanner := bufio.NewScanner(f)
	filescanner.Split(bufio.ScanLines)
	var filelines []string
	for filescanner.Scan() {
		filelines = append(filelines, filescanner.Text())
	}
	mentries := make(map[string]*Entry)
	for _, item := range filelines {
		//fmt.Printf("%v\n", item)
		tligne := strings.Split(item, ",")
		//fmt.Printf("%v longueur: %v\n", item, len(tligne))
		if len(tligne) > 4 {
			mentries[tligne[0]] = &Entry{tligne[0], tligne[1], tligne[2], 0}
		}
	}
	return Entries{filename, mentries}

}
func formatemax(lg int, n, m string) (string, string) {

	ln := n
	lm := m
	ln = fmt.Sprintf("%-*s", lg, n)
	lm = fmt.Sprintf("%-*s", lg, m)
	return ln, lm
}
func bilan(entrieUn, entrieDeux Entries) []string {
	tab := make([]string, 0)
	//	fmt.Print(entrieDeux)
	// faire defiler entree1
	//formatter les noms
	longmax := len(entrieUn.File)
	if len(entrieDeux.File) > longmax {
		longmax = len(entrieDeux.File)
	}
	nom1, nom2 := formatemax(longmax, entrieUn.File, entrieDeux.File)
	for cle, value := range entrieUn.Mentries {
		if value.Etat == 0 {
			if val2, ok := entrieDeux.Mentries[cle]; ok {
				if (val2.Password != value.Password) || (val2.User != value.User) {
					tmp := fmt.Sprintf("difference:%v - %v , %v, %v\n", nom1, value.Name, value.User, value.Password)
					tab = append(tab, tmp)
					tmp = fmt.Sprintf("----------:%v - %v , %v, %v\n", nom2, val2.Name, val2.User, val2.Password)
					tab = append(tab, tmp)
					val2.Etat = 2
					value.Etat = 2
				}
			} else {
				tmp := fmt.Sprintf("absent de %v:%v", entrieDeux.File, cle)
				tab = append(tab, tmp)
				value.Etat = 1
			}
		}
	}
	return tab
}

func main() {
	//r := strings.NewReader(Inp)
	argslen := len(os.Args)
	var filenameUn, filenameDeux string
	if argslen > 2 {
		filenameUn = os.Args[1]
		filenameDeux = os.Args[2]
	} else {
		os.Exit(1)
	}
	tmpUn := lire_fichier(filenameUn)
	//fmt.Printf("%v", tmp.Mentries)

	tmpDeux := lire_fichier(filenameDeux)
	//fmt.Printf("%v", tmp.Mentries)
	res := bilan(tmpUn, tmpDeux)
	for _, item := range res {
		fmt.Printf("%v", item)
	}
	res = bilan(tmpDeux, tmpUn)
	for _, item2 := range res {
		fmt.Printf("%v", item2)
	}

}
