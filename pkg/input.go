package pkg

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Initialize global variables
func init() {
	RoomConnections = make(map[string][]string)
}

// OpenFileIfArgsValid validates the command-line arguments and opens the specified file
func OpenFileIfArgsValid(arguments []string) *os.File {
	if len(arguments) != 2 {
		fmt.Println("Usage: program <input_file>")
		os.Exit(1)
	}
	
	file, err := os.Open(arguments[1])
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	
	return file
}

// ValidateRoomName validates the format of a start or end room name
func ValidateRoomName(roomDetails string) string {
	parts := strings.Split(roomDetails, " ")
	if len(parts) != 3 {
		log.Fatal("invalid input for room: incorrect number of parts")
	}
	
	roomName := parts[0]
	if roomName[0] == 'L' {
		log.Fatal("room name cannot start with 'L'")
	}
	
	return roomName
}

// ParseRoom parses and validates the format of a standard room entry
func ParseRoom(roomDetails string) string {
	parts := strings.Split(roomDetails, " ")
	if len(parts) != 3 {
		log.Fatal("invalid room data: incorrect number of parts")
	}
	
	roomName := parts[0]
	if roomName[0] == 'L' {
		log.Fatal("room name cannot start with 'L'")
	}
	
	return roomName
}

// AddRoomConnection adds a connection between two rooms to the adjacency map
func AddRoomConnection(connectionDetails string) {
	rooms := strings.Split(connectionDetails, "-")
	if len(rooms) != 2 {
		log.Fatalf("Invalid room connection format: %s", connectionDetails)
	}
	
	// Ensure rooms exist in the map
	if _, exists := RoomConnections[rooms[0]]; !exists {
		RoomConnections[rooms[0]] = []string{}
	}
	if _, exists := RoomConnections[rooms[1]]; !exists {
		RoomConnections[rooms[1]] = []string{}
	}
	
	// Add bidirectional connections
	RoomConnections[rooms[0]] = append(RoomConnections[rooms[0]], rooms[1])
	RoomConnections[rooms[1]] = append(RoomConnections[rooms[1]], rooms[0])
}

// IsSpecialCommand checks if a line is a valid special command
func IsSpecialCommand(command string) bool {
	if strings.HasPrefix(command, "##") {
		if command[2:] != "start" && command[2:] != "end" {
			log.Fatal("invalid special command")
		}
		return true
	}
	return false
}

// ParseAntsCount validates and converts the number of ants from a string to an integer
func ParseAntsCount(antsString string) int {
	ants, err := strconv.Atoi(antsString)
	if err != nil || ants <= 0 {
		log.Fatal("invalid ants number: must be a positive integer")
	}
	return ants
}

// ProcessInputFile reads, validates, and processes the input file
func ProcessInputFile(inputFile *os.File) error {
	// Reset global variables
	AntsCount = 0
	StartRoom = ""
	EndRoom = ""
	RoomList = []string{}
	Graphoverview = []byte{}
	RoomConnections = make(map[string][]string)

	// Read file contents
	fileContent, err := ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Process each line
	for i := 0; i < len(fileContent); i++ {
		line := fileContent[i]

		// Skip pure comments
		if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}

		// Append to graph overview
		Graphoverview = append(Graphoverview, []byte(line)...)
		Graphoverview = append(Graphoverview, '\n')

		// Process first line as the number of ants
		if i == 0 {
			AntsCount = ParseAntsCount(line)
			continue
		}

		// Process special commands
		if line == "##start" {
			if i+1 >= len(fileContent) {
				log.Fatal("missing room after ##start command")
			}
			StartRoom = ValidateRoomName(fileContent[i+1])
			RoomList = append(RoomList, StartRoom)
			i++
			continue
		}

		if line == "##end" {
			if i+1 >= len(fileContent) {
				log.Fatal("missing room after ##end command")
			}
			EndRoom = ValidateRoomName(fileContent[i+1])
			RoomList = append(RoomList, EndRoom)
			i++
			continue
		}

		// Process room entries
		if strings.Contains(line, " ") {
			roomName := ParseRoom(line)
			RoomList = append(RoomList, roomName)
		} else if strings.Contains(line, "-") {
			// Process room connections
			AddRoomConnection(line)
		}
	}

	// Validate input requirements
	if AntsCount == 0 {
		log.Fatal("no ants specified")
	}
	if StartRoom == "" {
		log.Fatal("no start room specified")
	}
	if EndRoom == "" {
		log.Fatal("no end room specified")
	}
	if len(RoomConnections) == 0 {
		log.Fatal("no room connections specified")
	}

	Graphoverview = append(Graphoverview, '\n')
	return nil
}