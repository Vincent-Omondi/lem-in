// file_reader.go
package pkg

import (
	"bufio"
	"os"
)

func ReadFile(file *os.File) ([]string, error) {
	defer file.Close()
	var content []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content, scanner.Err()
}
