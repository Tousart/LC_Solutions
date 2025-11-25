package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	route := []int{0}
	dfs(&graph, &res, 0, &route)
	return res
}

func dfs(graph, res *[][]int, vertex int, route *[]int) {
	for _, nextVertex := range (*graph)[vertex] {
		(*route) = append((*route), nextVertex)
		dfs(graph, res, nextVertex, route)
		(*route) = (*route)[:len((*route))-1]
	}

	if (*route)[len(*route)-1] == len(*graph)-1 {
		fullRoute := make([]int, len(*route))
		copy(fullRoute, *route)
		(*res) = append((*res), fullRoute)
	}
}

func main() {
	graph := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}
	fmt.Println(allPathsSourceTarget(graph))

	graph1 := [][]int{
		{2},
		{},
		{1},
	}
	fmt.Println(allPathsSourceTarget(graph1))
}
