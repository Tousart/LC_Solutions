package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	res := []string{}
	pathElems := strings.Split(path, "/")

	for _, val := range pathElems {
		if val == "." || val == "" {
			continue
		} else if val == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, val)
		}
	}

	return "/" + strings.Join(res, "/")
}

func main() {
	path := "/../"
	fmt.Println(simplifyPath(path))
}
