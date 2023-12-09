package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	util "lemonlabs/src/synchro/utils"
	"log"
	"os"
	"path/filepath"
	"time"
)

const backend string = "pro"

type records struct {
	path    string
	name    string
	hash    string
	date    string
	backend string
}

func storage() (state string) {
	const file string = "./synchro.db"
	//fmt.Println(create)
	//if _, err := db.Exec(create); err != nil {
	//	return "error create"
	//}
	return "Ok"
}

func get_hash(r io.Reader) string {
	hasher := sha256.New()
	//mb,err := io.ReadAll(r)
	//if err != nil {
	//    log.Fatal(err)
	//}
	// mettre un io.copy
	_, err := io.Copy(hasher, r)
	//hasher.Write(mb)
	if err != nil {
		log.Fatal(err)
	}
	result := hex.EncodeToString(hasher.Sum(nil))
	//fmt.Printf("%v\n", result)
	return result
}

func main() {
	var nFlag = flag.Int("prof", 2, "profondeur du répertoire")
	var pathFlag = flag.String("path", "./", "Chemin complet du repertoire")
	var repoFlag = flag.String("repo", "pro", "Etiquette de l'instance")
	flag.Parse()
	flag.Usage()
	directory := *pathFlag
	repo := *repoFlag
	proof := *nFlag
	madate := time.Now()
	proof = proof + 1
	smadate := madate.Format("02-01-2006")
	base := util.LastPath(directory)
	var trecords []records
	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			var ressource io.Reader
			var errfile error
			if err != nil {
				return err
			}
			//			fmt.Println("ch:", path)
			if !info.IsDir() {
				ressource, errfile = os.Open(path)
				if errfile != nil {
					log.Fatalln(errfile)
				}
				sp, sf, prof := util.ParsePath(path, base)
				//fmt.Println(sp, sf)
				//fmt.Printf("%v => %v\n", path, get_hash(ressource))
				fmt.Println("PROF", prof)
				trecords = append(trecords, records{sp, sf, get_hash(ressource), smadate, repo})

				return nil
			}
			if err != nil {
				log.Println(err)
			}
			return nil
		})
	fmt.Println(err)
	fmt.Println(trecords)
}
