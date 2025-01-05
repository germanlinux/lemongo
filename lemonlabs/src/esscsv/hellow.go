package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

var Inp = `"Sample Entry #2","Michael321","12345","https://keepass.info/help/kb/testform.html",""
"pgadmin","german","german","",""
"microsoft","germanlinux@yahoo.fr","Flora2103","",""
"zvm","root","rootvmeg","10.37.246.13",""
"collabra","eric.german@dgfip.finances.gouv.fr","Germanli","https://collabora.dgfip.finances.rie.gouv.fr/",""
`

type Entry struct {
	Name     string `csv:"name"`
	User     string `csv:"user"`
	password string `csv:"password"`
}

func main() {
	r := strings.NewReader(Inp)

	// read csv values using csv.Reader
	csvReader := csv.NewReader(r)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	fmt.Printf("data: %T\n", data)
	for _, item := range data {
		fmt.Printf("%v %T\n", item, item)

	}
}