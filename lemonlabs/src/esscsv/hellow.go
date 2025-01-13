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
	Name string
	User string
	Etat int
}

func main() {
	//r := strings.NewReader(Inp)
	f, err := os.Open("filecsv.csv")
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()

	// read csv values using csv.Reader
	//csvReader := csv.NewReader(f)
	//data, err := csvReader.ReadAll()
	filescanner := bufio.NewScanner(f)
	filescanner.Split(bufio.ScanLines)
	var filelines []string
	for filescanner.Scan() {
		filelines = append(filelines, filescanner.Text())
	}
	entree := make([]string, 0, 0)
	for _, item := range filelines {
		//fmt.Printf("%v\n", item)
		tligne := strings.Split(item, ",")
		//fmt.Printf("%v longueur: %v\n", item, len(tligne))
		if len(tligne) > 4 {
			entree = append(entree, item)
		}
	}
	for _, lig := range entree {
		fmt.Printf("%v\n", lig)
	}
}
