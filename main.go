// main.go
package main

import (
	"log"
	"os"

	pkg "github.com/Vincent-Omondi/lem-in/pkg"
)

func main() {
	pkg.ProcessInputFile(pkg.OpenFileIfArgsValid(os.Args))
	validways := pkg.FindPaths()
	if len(validways) == 0 || pkg.AntsCount == 0 {
		log.Fatal("ERROR: invalid data format")
	}
	os.Stdout.Write(pkg.Graphoverview)
	pkg.DispatchAnts(validways)
}
