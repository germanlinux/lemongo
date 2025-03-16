package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func exploreDirectory(path string) error {
	// Ouvrir le répertoire
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// Parcourir chaque entrée du répertoire
	for _, entry := range dirEntries {
		fullPath := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			// Si c'est un répertoire, on l'affiche et on explore récursivement
			fmt.Printf("Répertoire: %s\n", fullPath)
			err := exploreDirectory(fullPath)
			if err != nil {
				return err
			}
		} else {
			// Si c'est un fichier, on l'affiche
			fmt.Printf("Fichier; %s\n", fullPath)
		}
	}

	return nil
}

func main() {
	// Vérifier qu'un argument a été passé
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <répertoire>")
		return
	}

	// Récupérer le répertoire de départ
	startDir := os.Args[1]

	// Explorer le répertoire de départ
	err := exploreDirectory(startDir)
	if err != nil {
		fmt.Printf("Erreur lors de l'exploration du répertoire: %s\n", err)
	}
}
