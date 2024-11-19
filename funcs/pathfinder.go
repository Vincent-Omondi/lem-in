package lem

import "fmt"

func Search() [][]string {
	aaa := [][]string{}
	for _, v := range Ways[Start] {
		eee, deadend := Bfs(v)
		if !deadend {
			aaa = append(aaa, eee)
		}
	}
	sort(aaa)
	return aaa
}

func sort(unsorted [][]string) [][]string {
	for i := 0; i < len(unsorted); i++ {
		for j := i + 1; j < len(unsorted); j++ {
			if len(unsorted[i]) >= len(unsorted[j]) {
				unsorted[i], unsorted[j] = unsorted[j], unsorted[i]
			}
		}
	}
	return unsorted
}

func Bfs(s string) ([]string, bool) {
	levels := [][]string{Ways[Start]}
	if s == End {
		levels = append(levels, []string{End})
		return findway(levels), false

	}
	visited := make(map[string]bool)
	visited[Start] = true
	visited[s] = true
	tovisit := []string{s}
	for i := 0; i < len(tovisit); i++ {
		levl := []string{}
		visiting := tovisit[i]
		visited[visiting] = true
		for _, v := range Ways[visiting] {
			if !visited[v] {
				tovisit = append(tovisit, v)
				visited[v] = true
				levl = append(levl, v)
			}
			if v == End {
				return findway(levels), false
			}

		}
		levels = append(levels, levl)
	}
	return nil, true
}

func findway(levels [][]string) []string {
	curent := End
	way := []string{curent}
	for i := len(levels) - 1; i >= 0; i-- {
		for _, v := range levels[i] {
			if exist(v, curent) {
				way = append(way, v)
				curent = v
			}
		}
	}
	fmt.Println(way)
	//closeways(way[1:])
	way = append(way, Start)
	return flip(way)
}

func flip(s []string) []string {
	r := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

func exist(s, v string) bool {
	for _, t := range Ways[s] {
		if t == v {
			return true
		}
	}
	return false
}

func closeways(way []string) {
	for _, room := range way {
		for _, v := range Ways[room] {
			i := -1
			for o, x := range Ways[v] {
				if x == room {
					fmt.Println(room)
					i = o
					break
				}
			}
			if i != -1 {
				Ways[v] = append(Ways[v][:i], Ways[v][i+1:]...)
			}
		}
	}
}
