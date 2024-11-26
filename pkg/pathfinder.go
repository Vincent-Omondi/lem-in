package pkg

import "fmt"

func FindPaths() [][]string {
	var validPaths [][]string
	for _, neighbor := range RoomConnections[StartRoom] {
		path, isDeadEnd := TraverseGraph(neighbor)
		if !isDeadEnd {
			validPaths = append(validPaths, path)
		}
	}
	return SortPathsByLength(validPaths)
}

func SortPathsByLength(paths [][]string) [][]string {
	for i := 0; i < len(paths); i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) >= len(paths[j]) {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
	return paths
}

func TraverseGraph(startNode string) ([]string, bool) {
	levels := [][]string{RoomConnections[StartRoom]}
	if startNode == EndRoom {
		levels = append(levels, []string{EndRoom})
		return BuildPath(levels), false
	}

	visitedRooms := make(map[string]bool)
	visitedRooms[StartRoom] = true
	visitedRooms[startNode] = true
	nodesToVisit := []string{startNode}

	for i := 0; i < len(nodesToVisit); i++ {
		currentLevel := []string{}
		currentNode := nodesToVisit[i]
		visitedRooms[currentNode] = true

		for _, neighbor := range RoomConnections[currentNode] {
			if !visitedRooms[neighbor] {
				nodesToVisit = append(nodesToVisit, neighbor)
				visitedRooms[neighbor] = true
				currentLevel = append(currentLevel, neighbor)
			}
			if neighbor == EndRoom {
				return BuildPath(levels), false
			}
		}
		levels = append(levels, currentLevel)
	}
	return nil, true
}

func BuildPath(levels [][]string) []string {
	currentRoom := EndRoom
	path := []string{currentRoom}

	for i := len(levels) - 1; i >= 0; i-- {
		for _, room := range levels[i] {
			if IsConnected(room, currentRoom) {
				path = append(path, room)
				currentRoom = room
			}
		}
	}

	fmt.Println(path)
	path = append(path, StartRoom)
	return ReversePath(path)
}

func ReversePath(path []string) []string {
	reversed := []string{}
	for i := len(path) - 1; i >= 0; i-- {
		reversed = append(reversed, path[i])
	}
	return reversed
}

func IsConnected(roomA, roomB string) bool {
	for _, neighbor := range RoomConnections[roomA] {
		if neighbor == roomB {
			return true
		}
	}
	return false
}

// Uncomment and use this function if required to close paths after processing
// func ClosePaths(path []string) {
// 	for _, room := range path {
// 		for _, neighbor := range RoomConnections[room] {
// 			index := -1
// 			for i, connectedRoom := range RoomConnections[neighbor] {
// 				if connectedRoom == room {
// 					index = i
// 					break
// 				}
// 			}
// 			if index != -1 {
// 				RoomConnections[neighbor] = append(RoomConnections[neighbor][:index], RoomConnections[neighbor][index+1:]...)
// 			}
// 		}
// 	}
// }
