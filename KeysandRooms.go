package Leetcode

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool,len(rooms))
	dfs(rooms,rooms[0],visited)
	fmt.Println(visited)
	for _,data := range visited {
		if !data {
			return false
		}
	}

	return true
}

func dfs(room [][]int,rec []int,visited []bool) {
	for _, idx := range rec {
		if visited[idx] {
			continue
		}
		visited[idx] =true
		dfs(room,room[idx],visited)
	}
}

func main() {
	canVisitAllRooms([][]int{{1},{2},{3},{}})
}
