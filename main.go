package main

import (
	"log"
	"os"

	lem "github.com/Vincent-Omondi/lem-in/pkg"
)

func main() {
	lem.ValidData(lem.ValidArgs(os.Args))
	validways := lem.Search()
	if len(validways) == 0 || lem.Ants == 0 {
		log.Fatal("ERROR: invalid data format")
	}
	os.Stdout.Write(lem.Graphoverview)
	lem.Sendants(validways)
}
