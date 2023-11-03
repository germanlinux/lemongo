package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type chemin struct {
 	path  string
    description string
	tag string
}

type artefact struct {
	sujet string
	path  string
	methods []string
}
const entete_d2 = `
vars: {
	primaryColors: {
	  Accueil: {
		active: "#4baae5"
		border: black
	  }
	  Retour: {
		active: tomato
		border: black
	  }
	}
  }
  naviguation: {
	near: bottom-center
	Accueil: {
	  width: 100
	  height: 40
	  style: {
		border-radius: 5
		fill: ${primaryColors.Accueil.active}
		stroke: ${primaryColors.Accueil.border}
	  }
	}
	Accueil.link: "/"
	Retour: {
	  width: 100
	  height: 40
	  style: {
		border-radius: 5
		fill: ${primaryColors.Retour.active}
		stroke: ${primaryColors.Retour.border}
	  }
	}
	Retour.link: "D2_IAM.svg"
  `



func main() {
	myre := regexp.MustCompile(`\[(.+)\]`)
	myfile, err := ioutil.ReadFile("items.yaml")
	if err != nil {
		log.Fatal(err)
	}
    var recap  []artefact
	data := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal(myfile, &data)
	if err2 != nil {
	}
	// get title
	mtitre := data["info"].(map[string]interface{})
	titre := mtitre["title"]
    descr := mtitre["description"]
	paths := data["paths"]
	pathmap := paths.(map[string]interface{})
	//mapofpath := make(map[string]chemin)
	for  mypath, v := range pathmap {
			//fmt.Printf("%s -> %d\n", ke, v)
			methodmap := v.(map[string]interface{})
			keys := make([]string, len(methodmap))
			i := 0
			arte :=""
			for k := range methodmap {
    			keys[i] = k
				role := methodmap[k] 
				mrole := role.(map[string]interface{})
				arte = fmt.Sprintf("%v", mrole["tags"])

				//fmt.Printf("DEBUG %v %T", mrole["tags"],  mrole["tags"])
    			i++
			}
			/*
			fmt.Println("------------------------")
			fmt.Printf("artefact %v  path: %v :  keys : %v\n",arte,  ke,  keys)	
			fmt.Println("------------------------")
			*/
			matches := myre.FindStringSubmatch(arte)
			//fmt.Printf("DEBUG %v", matches[1]) 
			arte = matches[1]
			recap= append(recap,artefact{arte,mypath, keys } )
		}
	fmt.Printf("title:|md # %v\n## %v|" + "{near: top-center}\n", titre, descr)
	fmt.Println(entete_d2)
	dejavu := make(map[string]bool)
	for _, val := range(recap){
		if  !dejavu[val.sujet] {
			if len(val.methods) > 0 {  // on ferme la sequence D2 precedente
				fmt.Println("}")
			}
			fmt.Printf("%v:{\n",val.sujet )
			fmt.Println("shape: sql_table")
			dejavu[val.sujet] = true
			methodes := "\"["
			
			for _,item:= range(val.methods){
				methodes += strings.ToUpper(item) + " ,"
			}
			methodes = methodes[0:len(methodes) -1]
			methodes = methodes + "]"
			fmt.Printf("%v %v }\"\n",methodes, val.path)
		}
	}
	fmt.Println("}")		
	//fmt.Println(recap) 
	}
