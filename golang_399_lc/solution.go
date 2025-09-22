package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	powers := make(map[string]map[string]float64)
	for i, ab := range equations {
		a := ab[0]
		b := ab[1]

		if _, ok := powers[a]; !ok {
			powers[a] = make(map[string]float64)
		}

		if _, ok := powers[b]; !ok {
			powers[b] = make(map[string]float64)
		}

		powers[a][b] = values[i]
		powers[b][a] = 1 / values[i]
	}

	res := make([]float64, len(queries))

	for i, cd := range queries {
		c := cd[0]
		d := cd[1]
		if powers[c] == nil || powers[d] == nil {
			res[i] = -1.0
			continue
		} else if c == d {
			res[i] = 1.0
			continue
		}

		visited := make(map[string]bool)
		res[i] = 1 / dfs(powers, visited, c, d, 1.0)
	}

	return res
}

func dfs(powers map[string]map[string]float64, visited map[string]bool, target, current string, divider float64) float64 {
	if current == target {
		return divider
	}

	visited[current] = true

	for num, power := range powers[current] {
		if visited[num] {
			continue
		}

		result := dfs(powers, visited, target, num, divider*power)
		if result != -1.0 {
			return result
		}
	}

	visited[current] = false

	return -1.0
}

func main() {
	equations := [][]string{
		{"a", "b"},
		{"b", "c"},
	}
	values := []float64{2.0, 3.0}
	queries := [][]string{
		{"a", "c"},
		{"b", "a"},
		{"a", "e"},
		{"a", "a"},
		{"x", "x"},
	}

	fmt.Println(calcEquation(equations, values, queries))

	equations1 := [][]string{
		{"a", "b"},
		{"c", "d"},
	}
	values1 := []float64{1.0, 1.0}
	queries1 := [][]string{
		{"a", "c"},
		{"b", "d"},
		{"b", "a"},
		{"d", "c"},
	}

	fmt.Println(calcEquation(equations1, values1, queries1))
}
