package lem

import (
    "log"
    "os"
    "strconv"
    "strings"
)

// InputProcessor handles file reading and data validation
type InputProcessor struct {
    graphState *GraphState
    overview   []byte
}

// NewInputProcessor creates a new input processor instance
func NewInputProcessor(state *GraphState) *InputProcessor {
    return &InputProcessor{
        graphState: state,
        overview:   make([]byte, 0, 1024), // Reasonable initial capacity
    }
}

// OpenInputFile validates command line arguments and opens the input file
func OpenInputFile(args []string) (*os.File, error) {
    if len(args) != 2 {
        return nil, nil
    }
    
    return os.Open(args[1])
}

// ProcessInputFile handles the complete input processing workflow
func (p *InputProcessor) ProcessInputFile(file *os.File) error {
    fileLines, err := ReadFile(file)
    if err != nil {
        return err
    }

    return p.processFileContents(fileLines)
}

// processFileContents validates and processes each line of input
func (p *InputProcessor) processFileContents(lines []string) error {
    for i := 0; i < len(lines); i++ {
        currentLine := lines[i]
        
        // Append to overview
        p.appendToOverview(currentLine)

        // Skip regular comments
        if p.isRegularComment(currentLine) {
            continue
        }

        // Process different line types
        if i == 0 {
            p.graphState.EntityCount = p.parseAntsCount(currentLine)
            continue
        }

        // Handle special commands (##start, ##end)
        if isCommand := p.processCommand(currentLine, &i, lines); isCommand {
            continue
        }

        // Process room or connection
        p.processRoomOrConnection(currentLine)
    }

    // Validate final state
    if p.graphState.StartRoom == "" || p.graphState.EndRoom == "" {
        log.Fatal("end or start room missing")
    }

    p.overview = append(p.overview, '\n')
    return nil
}

// parseAntsCount validates and converts the ants number
func (p *InputProcessor) parseAntsCount(value string) int {
    count, err := strconv.Atoi(value)
    if err != nil {
        log.Fatal("invalid ants number")
    }
    return count
}

// validateRoomFormat checks room format and name constraints
func (p *InputProcessor) validateRoomFormat(roomData string) string {
    parts := strings.Split(roomData, " ")
    if len(parts) != 3 {
        log.Fatal("invalid input for room")
    }

    roomName := parts[0]
    if roomName[0] == 'L' {
        log.Fatal("room name cant start with 'L'")
    }
    return roomName
}

// processConnection adds bidirectional connection between rooms
func (p *InputProcessor) processConnection(connection string) {
    rooms := strings.Split(connection, "-")
    p.graphState.Connections[rooms[0]] = append(p.graphState.Connections[rooms[0]], rooms[1])
    p.graphState.Connections[rooms[1]] = append(p.graphState.Connections[rooms[1]], rooms[0])
}

// isRegularComment checks if line is a regular comment
func (p *InputProcessor) isRegularComment(line string) bool {
    return strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##")
}

// processCommand handles ##start and ##end commands
func (p *InputProcessor) processCommand(line string, lineIndex *int, lines []string) bool {
    if !strings.HasPrefix(line, "##") {
        return false
    }

    command := line[2:]
    if command != "start" && command != "end" {
        log.Fatal("invalid command")
    }

    // Validate next line exists
    if *lineIndex == len(lines)-1 {
        log.Fatal("end or start room missing")
    }

    // Process the room on the next line
    nextLine := lines[*lineIndex+1]
    roomName := p.validateRoomFormat(nextLine)
    
    // Update state based on command
    if command == "start" {
        p.graphState.StartRoom = roomName
    } else {
        p.graphState.EndRoom = roomName
    }

    p.graphState.RoomList = append(p.graphState.RoomList, roomName)
    *lineIndex++
    
    return true
}

// processRoomOrConnection handles regular room definitions or connections
func (p *InputProcessor) processRoomOrConnection(line string) {
    if strings.Contains(line, " ") {
        roomName := p.validateRoomFormat(line)
        p.graphState.RoomList = append(p.graphState.RoomList, roomName)
    } else if strings.Contains(line, "-") {
        p.processConnection(line)
    }
}

// appendToOverview adds a line to the graph overview
func (p *InputProcessor) appendToOverview(line string) {
    p.overview = append(p.overview, []byte(line)...)
    p.overview = append(p.overview, '\n')
}