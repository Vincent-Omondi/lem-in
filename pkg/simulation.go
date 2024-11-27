// simulator.go
package pkg

import (
	"fmt"
	"strconv"
)

func DispatchAnts(paths [][]string) {
	antGroups := [][]string{}
	antID := 1
	if len(paths) > AntsCount {
		paths = paths[:AntsCount]
	}
	for i := 0; i < len(paths); i++ {
		antGroup := []string{}
		for j := 0; j < AntsCount/len(paths); j++ {
			if antID > AntsCount {
				break
			}
			antGroup = append(antGroup, "L"+strconv.Itoa(antID))
			antID++
		}
		if i == 0 && antID <= AntsCount {
			antGroup = append(antGroup, "L"+strconv.Itoa(antID))
			antID++
		}
		antGroups = append(antGroups, antGroup)
	}
	if antID <= AntsCount {
		for i := 0; i < len(antGroups); i++ {
			if antID > AntsCount {
				break
			}
			antGroups[i] = append(antGroups[i], "L"+strconv.Itoa(antID))
			antID++
		}
	}
	ControlTraffic(antGroups, paths)
}

func ControlTraffic(antGroups, paths [][]string) {
	traffic := make(map[string]int)
	unavailableRooms := make(map[string]bool)
	completedAnts := []string{}
	for len(completedAnts) != AntsCount {
		for i := 0; i < len(paths); i++ {
			unavailableRooms[EndRoom] = false
			for s := 0; s < len(antGroups[i]); s++ {
				ant := antGroups[i][s]
				if !unavailableRooms[paths[i][traffic[ant]+1]] {
					if paths[i][traffic[ant]+1] == EndRoom {
						unavailableRooms[paths[i][traffic[ant]]] = false
						completedAnts = append(completedAnts, ant)
						delete(traffic, ant)
						antGroups[i] = append(antGroups[i][:s], antGroups[i][s+1:]...)
						fmt.Printf("%v-%v ", ant, EndRoom)
						s--
						unavailableRooms[EndRoom] = true
						continue
					} else {
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

func SortAnts(ants []string) {
	for i := 0; i < len(ants); i++ {
		for j := i + 1; j < len(ants); j++ {
			if ants[j] < ants[i] {
				ants[j], ants[i] = ants[i], ants[j]
			}
		}
	}
}
