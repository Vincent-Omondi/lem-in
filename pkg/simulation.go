package lem

import (
	"fmt"
	"strconv"
)

func Sendants(ways [][]string) {
	antgroups := [][]string{}
	antid := 1
	for i := 0; i < len(ways); i++ {
		antgroup := []string{}
		for j := 0; j < Ants/len(ways); j++ {
			if antid > Ants {
				break
			}
			antgroup = append(antgroup, "L"+strconv.Itoa(antid))
			antid++

		}
		if i == 0 && antid < Ants {
			antgroup = append(antgroup, "L"+strconv.Itoa(antid))
			antid++
		}
		antgroups = append(antgroups, antgroup)
	}
	controltrafic(antgroups, ways)
}

func controltrafic(antgroups, ways [][]string) {
	trafic := make(map[string]int)
	unavailablerooms := make(map[string]bool)
	finished := []string{}
	for len(finished) != Ants {
		for i := 0; i < len(ways); i++ {
			unavailablerooms[End] = false
			for s := 0; s < len(antgroups[i]); s++ {
				ant := antgroups[i][s]
				if !unavailablerooms[ways[i][trafic[ant]+1]] {
					if ways[i][trafic[ant]+1] == End {
						unavailablerooms[ways[i][trafic[ant]]] = false
						finished = append(finished, ant)
						delete(trafic, ant)
						antgroups[i] = append(antgroups[i][:s], antgroups[i][s+1:]...)
						fmt.Printf("%v-%v ", ant, End)
						s--
						unavailablerooms[End] = true
						continue
					} else {
						fmt.Printf("%v-%v ", ant, ways[i][trafic[ant]+1])
						unavailablerooms[ways[i][trafic[ant]+1]] = true
						unavailablerooms[ways[i][trafic[ant]]] = false
						trafic[ant]++
					}
				}
			}
		}
		fmt.Println()
	}
}
