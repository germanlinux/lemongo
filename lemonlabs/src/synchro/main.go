package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	util "lemonlabs/src/synchro/utils"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const create string = `
  CREATE TABLE IF NOT EXISTS fingerprint (
  path TEXT ,
  name  TEXT,
  hash TEXT, 
  date TEXT,
  backend TEXT ,
  PRIMARY KEY (path, name)
   );`
const backend string = "pro"

func storage() (state string) {
	const file string = "./synchro.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		fmt.Printf("err %v\n", err)
		panic("stop")
	}
	fmt.Println(create)
	if _, err := db.Exec(create); err != nil {
		return "error create"
	}
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
	//flag.Usage()
	directory := *pathFlag
	repo := *repoFlag
	proof := *nFlag
	fmt.Println(repo, proof)
	if len(os.Args) > 1 {
		directory = os.Args[1]
		// do something with command
	}
	base := util.LastPath(directory)
	fmt.Println("base:", base)
	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			var ressource io.Reader
			var errfile error
			if err != nil {
				return err
			}
			fmt.Println("ch:", path)
			if !info.IsDir() {
				ressource, errfile = os.Open(path)
				if errfile != nil {
					log.Fatalln(errfile)
				}
				sp, sf := util.ParsePath(path, base)
				fmt.Println(sp, sf)
				fmt.Printf("%v => %v\n", path, get_hash(ressource))
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(storage())

}
