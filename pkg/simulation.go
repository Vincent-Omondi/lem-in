package pkg

import (
	"fmt"
	"strconv"
)

func DispatchAnts(paths [][]string) {
	antGroups := [][]string{}
	antID := 1

	// Distribute ants among the paths
	for i := 0; i < len(paths); i++ {
		antGroup := []string{}
		// Distribute ants equally among the paths
		for j := 0; j < AntsCount/len(paths); j++ {
			if antID > AntsCount {
				break
			}
			antGroup = append(antGroup, "L"+strconv.Itoa(antID))
			antID++
		}

		// If there are remaining ants, assign them to the first path
		if i == 0 && antID <= AntsCount {
			antGroup = append(antGroup, "L"+strconv.Itoa(antID))
			antID++
		}
		antGroups = append(antGroups, antGroup)
	}
	ControlTraffic(antGroups, paths)
}

func ControlTraffic(antGroups, paths [][]string) {
	traffic := make(map[string]int)          // Tracks the current position of each ant
	unavailableRooms := make(map[string]bool) // Tracks which rooms are unavailable
	completedAnts := []string{}              // Tracks completed ants

	// Continue until all ants have reached the end
	for len(completedAnts) != AntsCount {
		for i := 0; i < len(paths); i++ {
			unavailableRooms[EndRoom] = false

			// Move each ant along its path
			for s := 0; s < len(antGroups[i]); s++ {
				ant := antGroups[i][s]

				// Check if the next room in the path is available
				if !unavailableRooms[paths[i][traffic[ant]+1]] {
					if paths[i][traffic[ant]+1] == EndRoom {
						// If the ant reaches the end, mark it as finished
						unavailableRooms[paths[i][traffic[ant]]] = false
						completedAnts = append(completedAnts, ant)
						delete(traffic, ant)
						antGroups[i] = append(antGroups[i][:s], antGroups[i][s+1:]...)
						fmt.Printf("%v-%v ", ant, EndRoom)
						s-- // Adjust the index since we removed an ant from the group
						unavailableRooms[EndRoom] = true
						continue
					} else {
						// Move the ant to the next room in the path
						fmt.Printf("%v-%v ", ant, paths[i][traffic[ant]+1])
						unavailableRooms[paths[i][traffic[ant]+1]] = true
						unavailableRooms[paths[i][traffic[ant]]] = false
						traffic[ant]++
					}
				}
			}
		}
		fmt.Println()
	}
}
