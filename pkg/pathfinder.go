// pathfinder.go
package pkg

import (
	"fmt"
	"sort"
)

func FindPaths() [][]string {
	DepthFirstSearch(StartRoom)
	SortPathsByLength(ValidPaths)
	SetPathRatings()
	RatePaths()
	return CombinePaths(len(RoomConnections[StartRoom]))
}

func SetPathRatings() {
	for index := range ValidPaths {
		PathRatings[index] = len(ValidPaths[index])
	}
}

func ArePathsIdentical(path1, path2 []string) bool {
	for i, room := range path1 {
		if path2[i] != room {
			return false
		}
	}
	return true
}

func HasOverlap(paths ...[]string) bool {
	for _, path := range paths {
		for _, otherPath := range paths {
			if CheckConflict(otherPath, path) {
				return true
			}
		}
	}
	return false
}

func CheckConflict(path1, path2 []string) bool {
	for i := 1; i < len(path1)-1; i++ {
		room1 := path1[i]
		for j := 1; j < len(path2)-1; j++ {
			room2 := path2[j]
			if !ArePathsIdentical(path1, path2) && room1 == room2 {
				return true
			}
		}
	}
	return false
}

func RatePaths() {
	for i, path1 := range ValidPaths {
		for j, path2 := range ValidPaths {
			if i != j && CheckConflict(path2, path1) {
				PathRatings[i]++
			}
		}
	}
}

func SortPathInfo(pathInfo []PathInfo) []PathInfo {
	for i := 0; i < len(pathInfo)-1; i++ {
		for j := i + 1; j < len(pathInfo); j++ {
			if pathInfo[i].Rating > pathInfo[j].Rating {
				pathInfo[i], pathInfo[j] = pathInfo[j], pathInfo[i]
			}
		}
	}
	return pathInfo
}

func GeneratePathInfo() []PathInfo {
	info := []PathInfo{}
	for index, rating := range PathRatings {
		info = append(info, PathInfo{
			Rating: rating,
			Index:  index,
		})
	}
	return SortPathInfo(info)
}

func CombinePaths(maxPaths int) [][]string {
	selectedPaths := [][]string{}
	pathInfo := GeneratePathInfo()
	maxIndex := 0

	for u := 0; u < len(pathInfo); u++ {
		entry := pathInfo[u]
		index := entry.Index
		if len(ValidPaths[index]) > maxIndex {
			maxIndex = index
		}
		temp := append(selectedPaths, ValidPaths[index])
		if !HasOverlap(temp...) {
			selectedPaths = append(selectedPaths, ValidPaths[index])
		} else if len(ValidPaths[maxIndex]) > 2*len(ValidPaths[index]) {
			maxIndex = index
			u = 0
		}
		SortPathsByLength(selectedPaths)
		if len(selectedPaths) == maxPaths {
			break
		}
	}
	return selectedPaths
}

func SortPathsByLength(paths [][]string) [][]string {
	for i := 0; i < len(paths); i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) > len(paths[j]) {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
	return paths
}

func DepthFirstSearch(currentRoom string) {
	TraversalStack = append(TraversalStack, currentRoom)

	if VisitedRooms[currentRoom] {
		TraversalStack = TraversalStack[:len(TraversalStack)-1]
		return
	}

	if currentRoom == EndRoom {
		path := []string{}
		path = append(path, TraversalStack...)
		ValidPaths = append(ValidPaths, path)
		TraversalStack = TraversalStack[:len(TraversalStack)-1]
		return
	}
	VisitedRooms[currentRoom] = true
	for _, neighbor := range RoomConnections[currentRoom] {
		DepthFirstSearch(neighbor)
	}
	TraversalStack = TraversalStack[:len(TraversalStack)-1]
	VisitedRooms[currentRoom] = false
}

// Additional pathfinding logic

func SearchMax() [][]string {
	VisitedRooms[StartRoom] = true
	for i := 0; i < len(RoomConnections[StartRoom]); i++ {
		node := RoomConnections[StartRoom][i]
		TraverseGraph(node)
		fmt.Println(ValidPaths)
	}
	sortSolutions()
	return ValidPaths
}

func sortSolutions() {
	sort.Slice(ValidPaths, func(i, j int) bool {
		// First, sort by the length of the solution
		if len(ValidPaths[i]) != len(ValidPaths[j]) {
			return len(ValidPaths[i]) < len(ValidPaths[j])
		}
		// If lengths are the same, use the compare function to decide order
		return compare(ValidPaths[i][1:len(ValidPaths[i])-1], ValidPaths[j][1:len(ValidPaths[j])-1])
	})
}

func compare(s1, s2 []string) bool {
	if len(s1) != len(s2) || len(s1) == 0 || len(s2) == 0 {
		return false
	}

	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}
	return true
}

func TraverseGraph(startNode string) bool {
	parent := make(map[string]string)
	parent[startNode] = StartRoom

	if startNode == EndRoom {
		ValidPaths = append(ValidPaths, BuildPath(parent))
		return false
	}

	VisitedRooms[StartRoom] = true
	queue := []string{startNode}
	VisitedRooms[startNode] = true

	for i := 0; i < len(queue); i++ {
		currentNode := queue[i]
		for _, neighbor := range RoomConnections[currentNode] {
			if !VisitedRooms[neighbor] {
				VisitedRooms[neighbor] = true
				parent[neighbor] = currentNode
				queue = append(queue, neighbor)
			}
			if neighbor == EndRoom {
				ValidPaths = append(ValidPaths, BuildPath(parent))
				ClosePaths()
				return false
			}
		}
	}
	return true
}

func BuildPath(parent map[string]string) []string {
	current := EndRoom
	VisitedRooms = make(map[string]bool)
	path := []string{current}
	for current != StartRoom {
		path = append(path, parent[current])
		current = parent[current]
	}
	return ReversePath(path)
}

func ClosePaths() {
	VisitedRooms = make(map[string]bool)
	for _, path := range ValidPaths {
		for _, room := range path[1 : len(path)-1] {
			VisitedRooms[room] = true
		}
	}
}

func ReversePath(path []string) []string {
	reversed := []string{}
	for i := len(path) - 1; i >= 0; i-- {
		reversed = append(reversed, path[i])
	}
	return reversed
}
