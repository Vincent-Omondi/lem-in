package pkg

import (
	"bufio"
	"os"
)

func ReadFile(inputFile *os.File) ([]string, error) {
	defer inputFile.Close()
	var fileContent []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}
	return fileContent, scanner.Err()
}
