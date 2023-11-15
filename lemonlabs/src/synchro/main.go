package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func get_hash(r  io.Reader) string {
	hasher := sha256.New()
    //mb,err := io.ReadAll(r)
	//if err != nil {
    //    log.Fatal(err)
	//}	
// mettre un io.copy
    _,err := io.Copy(hasher, r)
    //hasher.Write(mb)
	if err != nil {
		   log.Fatal(err)
		}
	result := hex.EncodeToString(hasher.Sum(nil))
	//fmt.Printf("%v\n", result)
	return result
}

func main() {
	
    err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			var ressource io.Reader
			var errfile error
			if err != nil {
				return err
			}
		//fmt.Println(path, info.Size())
		if !info.IsDir() {
			ressource, errfile = os.Open(path) 
			if errfile != nil {
				log.Fatalln(errfile)
			}
			fmt.Printf("%v => %v\n", path, get_hash(ressource))
		}
		return nil
		})
	if err != nil {
	log.Println(err)
	}
/*
    file :=os.Args[1]
	h:= get_hash(file)
	fmt.Println(h)
    //os.Stdout.WriteString(result)
*/
}

