// input.go
package pkg

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func OpenFileIfArgsValid(arguments []string) (*os.File, error) {
	if len(arguments) != 2 {
		return nil, logError("invalid number of arguments")
	}
	inputFile, err := os.Open(arguments[1])
	if err != nil {
		return nil, err
	}
	return inputFile, nil
}

func ProcessInputFile(inputFile *os.File) (string, error) {
	fileContent, err := ReadFile(inputFile)
	if err != nil {
		return "", err
	}
	counter := 0
	for i := 0; i < len(fileContent); i++ {
		line := fileContent[i]
		if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		} else if strings.HasPrefix(line, "##") {
			if line[2:] != "start" && line[2:] != "end" {
				return "", logError("invalid special command")
			}
		}
		Graphoverview = append(Graphoverview, fileContent[i]...)
		Graphoverview = append(Graphoverview, '\n')
		if i == 0 {
			AntsCount, err = strconv.Atoi(line)
			if err != nil {
				return "", logError("invalid number of ants")
			}
			continue
		}

		if line == "##start" {
			counter++
			if i == len(fileContent)-1 {
				return "", logError("missing room after ##start")
			}
			StartRoom = strings.Split(fileContent[i+1], " ")[0]
			if len(strings.Split(fileContent[i+1], " ")) != 3 {
				return "", logError("invalid input for start room")
			}
			if StartRoom[0] == 'L' {
				return "", logError("room name cannot start with 'L'")
			}
			Graphoverview = append(Graphoverview, fileContent[i+1]...)
			Graphoverview = append(Graphoverview, '\n')
			RoomList = append(RoomList, StartRoom)
			i++
		} else if line == "##end" {
			counter++
			if i == len(fileContent)-1 {
				return "", logError("missing room after ##end")
			}
			EndRoom = strings.Split(fileContent[i+1], " ")[0]
			if len(strings.Split(fileContent[i+1], " ")) != 3 {
				return "", logError("invalid input for end room")
			}
			if EndRoom[0] == 'L' {
				return "", logError("room name cannot start with 'L'")
			}
			RoomList = append(RoomList, EndRoom)
			Graphoverview = append(Graphoverview, fileContent[i+1]...)
			Graphoverview = append(Graphoverview, '\n')
			i++
		} else if strings.Contains(line, " ") {
			counter++
			roomName := strings.Split(line, " ")[0]
			if len(strings.Split(line, " ")) != 3 {
				return "", logError("invalid room data")
			}
			if roomName[0] == 'L' {
				return "", logError("room name cannot start with 'L'")
			}
			RoomList = append(RoomList, roomName)
		} else {
			room1 := strings.Split(line, "-")[0]
			room2 := strings.Split(line, "-")[1]
			if slices.Contains(RoomConnections[room1], room2) {
				return "", logError("duplicate connection between rooms")
			}
			counter += 2
			RoomConnections[room1] = append(RoomConnections[room1], room2)
			RoomConnections[room2] = append(RoomConnections[room2], room1)
		}
	}
	if StartRoom == "" || EndRoom == "" {
		return "", logError("start or end room missing")
	}
	searchMethod := "dfs"
	if counter > 100 {
		searchMethod = "bfs"
	}
	Graphoverview = append(Graphoverview, '\n')
	return searchMethod, nil
}

func logError(message string) error {
	log.Println(message)
	return fmt.Errorf(message)
}
