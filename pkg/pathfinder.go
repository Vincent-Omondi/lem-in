package lem

const (
    // DefaultCapacity is the initial capacity for various collections
    DefaultCapacity = 10
)

// GraphState represents the current state of the maze/graph
type GraphState struct {
    // Adjacency list representing connections between rooms
    Connections map[string][]string
    // Map to track empty rooms
    EmptyRooms map[string]bool
    // List of all rooms in the graph
    RoomList []string
    // Start and end points in the maze
    StartRoom, EndRoom string
    // Number of entities to route through the maze
    EntityCount int
    // Raw graph data for debugging/visualization
    GraphData []byte
}

// NewGraphState initializes a new GraphState with default values
func NewGraphState() *GraphState {
    return &GraphState{
        Connections: make(map[string][]string),
        EmptyRooms:  make(map[string]bool),
        RoomList:    make([]string, 0, DefaultCapacity),
    }
}

// FindAllPaths returns all possible paths from start to end room
func (g *GraphState) FindAllPaths() [][]string {
    validPaths := [][]string{}

    // Try each possible first step from the start room
    for _, nextRoom := range g.Connections[g.StartRoom] {
        path, isDeadEnd := g.breadthFirstSearch(nextRoom)
        if !isDeadEnd {
            validPaths = append(validPaths, path)
        }
    }

    return g.sortPathsByLength(validPaths)
}

// sortPathsByLength sorts paths by their length in ascending order
func (g *GraphState) sortPathsByLength(paths [][]string) [][]string {
    for i := 0; i < len(paths); i++ {
        for j := i + 1; j < len(paths); j++ {
            if len(paths[i]) >= len(paths[j]) {
                paths[i], paths[j] = paths[j], paths[i]
            }
        }
    }
    return paths
}

// breadthFirstSearch performs BFS from a given starting room
func (g *GraphState) breadthFirstSearch(startRoom string) ([]string, bool) {
    // Track visited rooms to avoid cycles
    visited := make(map[string]bool)
    visited[g.StartRoom] = true
    visited[startRoom] = true

    // Initialize BFS queue and levels
    toVisit := []string{startRoom}
    levels := [][]string{g.Connections[g.StartRoom]}

    // Special case: direct path to end
    if startRoom == g.EndRoom {
        levels = append(levels, []string{g.EndRoom})
        return g.reconstructPath(levels), false
    }

    // Standard BFS implementation
    for i := 0; i < len(toVisit); i++ {
        currentLevel := []string{}
        currentRoom := toVisit[i]
        visited[currentRoom] = true

        for _, neighbor := range g.Connections[currentRoom] {
            if !visited[neighbor] {
                toVisit = append(toVisit, neighbor)
                visited[neighbor] = true
                currentLevel = append(currentLevel, neighbor)
            }

            if neighbor == g.EndRoom {
                levels = append(levels, currentLevel)
                return g.reconstructPath(levels), false
            }
        }

        levels = append(levels, currentLevel)
    }

    return nil, true
}

// reconstructPath builds the final path from BFS levels
func (g *GraphState) reconstructPath(levels [][]string) []string {
    currentRoom := g.EndRoom
    path := []string{currentRoom}

    // Traverse levels backwards to find the path
    for i := len(levels) - 1; i >= 0; i-- {
        for _, room := range levels[i] {
            if g.areRoomsConnected(room, currentRoom) {
                path = append(path, room)
                currentRoom = room
            }
        }
    }

    path = append(path, g.StartRoom)
    return g.reversePath(path)
}

// reversePath reverses the order of rooms in a path
func (g *GraphState) reversePath(path []string) []string {
    reversed := make([]string, len(path))
    for i := len(path) - 1; i >= 0; i-- {
        reversed[len(path)-1-i] = path[i]
    }
    return reversed
}

// areRoomsConnected checks if two rooms are directly connected
func (g *GraphState) areRoomsConnected(room1, room2 string) bool {
    for _, connectedRoom := range g.Connections[room1] {
        if connectedRoom == room2 {
            return true
        }
    }
    return false
}

// removeConnection removes a connection between two rooms
// func (g *GraphState) removeConnection(room1, room2 string) {
//     connections := g.Connections[room2]
//     for i, room := range connections {
//         if room == room1 {
//             // Remove the connection by slicing
//             g.Connections[room2] = append(connections[:i], connections[i+1:]...)
//             break
//         }
//     }
// }

// // closePathConnections removes connections along a given path
// func (g *GraphState) closePathConnections(path []string) {
//     for _, room := range path {
//         for _, neighbor := range g.Connections[room] {
//             g.removeConnection(room, neighbor)
//         }
//     }
// }