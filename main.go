package main

import (
	"log"
	"os"

	pkg "github.com/Vincent-Omondi/lem-in/pkg"
)

func main() {
	inputFile, err := pkg.OpenFileIfArgsValid(os.Args)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	searchMethod, err := pkg.ProcessInputFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to process input file: %v", err)
	}

	validPaths := [][]string{}
	switch searchMethod {
	case "bfs":
		validPaths = pkg.SearchMax()
	default:
		validPaths = pkg.FindPaths()
	}

	if len(validPaths) == 0 || pkg.AntsCount == 0 {
		log.Fatal("ERROR: invalid data format")
	}

	_, writeErr := os.Stdout.Write(pkg.Graphoverview)
	if writeErr != nil {
		log.Fatalf("Failed to write graph overview: %v", writeErr)
	}

	pkg.DispatchAnts(validPaths)
}
