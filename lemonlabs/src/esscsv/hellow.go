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

func bilan(entrieUn, entrieDeux Entries) int {
	//	fmt.Print(entrieDeux)
	// faire defiler entree1
	for cle, value := range entrieUn.Mentries {
		if value.Etat == 0 {
			if val2, ok := entrieDeux.Mentries[cle]; ok {
				if (val2.Password != value.Password) or (val2.User != value.User) {
					fmt.Printf("difference: %v , %v", value, val2)
					Value.Etat = 2 					
				} 
			} else {
				fmt.Println("absent:", entrieDeux.File, cle)
				Value.Etat = 1
			}
		}
	}
	return 0
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
	for cle, val := range tmpUn.Mentries {
		fmt.Printf("%v  -  %v\n", cle, val)
	}
	tmpDeux := lire_fichier(filenameDeux)
	//fmt.Printf("%v", tmp.Mentries)
	for cle, val := range tmpDeux.Mentries {
		fmt.Printf("%v  -  %v\n", cle, val)
	}
	toto := bilan(tmpUn, tmpDeux)
	fmt.Println(toto)
}
