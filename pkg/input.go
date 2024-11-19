package lem

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// ValidArgs checks if command line arguments are valid and opens the file
func ValidArgs(args []string) *os.File {
    if len(args) != 2 {
        return nil
    }
    data, err := os.Open(args[1])
    if err != nil {
        return nil
    }
    return data
}

// ValidateStartEndRoom validates start/end room format
func ValidateStartEndRoom(roomData string) string {
    room := strings.Split(roomData, " ")[0]
    if len(strings.Split(roomData, " ")) != 3 {
        log.Fatal("invalid input for room")
    }
    if room[0] == 'L' {
        log.Fatal("room name cant start with 'L'")
    }
    return room
}

// ProcessRoom handles regular room data
func ProcessRoom(roomData string) string {
    room := strings.Split(roomData, " ")[0]
    if len(strings.Split(roomData, " ")) != 3 {
        log.Fatal("invalid room data")
    }
    if room[0] == 'L' {
        log.Fatal("room name cant start with 'L'")
    }
    return room
}

// ProcessConnection adds connection between rooms to Ways map
func ProcessConnection(connection string, Ways map[string][]string) {
    parts := strings.Split(connection, "-")
    Ways[parts[0]] = append(Ways[parts[0]], parts[1])
    Ways[parts[1]] = append(Ways[parts[1]], parts[0])
}

// ValidateCommand validates special commands like ##start and ##end
func ValidateCommand(cmd string) bool {
    if strings.HasPrefix(cmd, "##") {
        if cmd[2:] != "start" && cmd[2:] != "end" {
            log.Fatal("invalid msg")
        }
        return true
    }
    return false
}

// ProcessAnts validates and converts ants number
func ProcessAnts(value string) int {
    ants, err := strconv.Atoi(value)
    if err != nil {
        log.Fatal("invalid ants number")
    }
    return ants
}

// ValidData processes and validates the input file
func ValidData(file *os.File) error {
    data, err := ReadFile(file)
    if err != nil {
        return err
    }

    for i := 0; i < len(data); i++ {
        v := data[i]
        
        // Skip comments
        if strings.HasPrefix(v, "#") && !strings.HasPrefix(v, "##") {
            continue
        }

        // Update overview
        Graphoverview = append(Graphoverview, data[i]...)
        Graphoverview = append(Graphoverview, '\n')

        // Process first line (ants number)
        if i == 0 {
            Ants = ProcessAnts(v)
            continue
        }

        // Process start command
        if v == "##start" {
            if i == len(data)-1 {
                log.Fatal("end or start room missing")
            }
            Start = ValidateStartEndRoom(data[i+1])
            Rooms = append(Rooms, Start)
            i++
            continue
        }

        // Process end command
        if v == "##end" {
            if i == len(data)-1 {
                log.Fatal("end or start room missing")
            }
            End = ValidateStartEndRoom(data[i+1])
            Rooms = append(Rooms, End)
            i++
            continue
        }

        // Process room or connection
        if strings.Contains(v, " ") {
            room := ProcessRoom(v)
            Rooms = append(Rooms, room)
        } else if strings.Contains(v, "-") {
            ProcessConnection(v, Ways)
        }
    }

    // Validate start and end rooms exist
    if Start == "" || End == "" {
        log.Fatal("end or start room missing")
    }

    Graphoverview = append(Graphoverview, '\n')
    return nil
}