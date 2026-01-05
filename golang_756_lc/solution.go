package main

import "fmt"

func pyramidTransition(bottom string, allowed []string) bool {
	invalidLayers := make(map[string]struct{})

	trios := make(map[string][]byte)
	for _, trio := range allowed {
		trios[trio[:2]] = append(trios[trio[:2]], trio[2])
	}

	return nextLayer(bottom, nil, 0, invalidLayers, trios)
}

func nextLayer(prevLayer string, currLayer []byte, idx int, invalidLayers map[string]struct{}, trios map[string][]byte) bool {
	if len(prevLayer) == 1 {
		return true
	}

	if idx == len(prevLayer)-1 {
		strCurrLayer := string(currLayer)

		if _, ok := invalidLayers[strCurrLayer]; ok {
			return false
		}

		next := nextLayer(strCurrLayer, nil, 0, invalidLayers, trios)
		if !next {
			invalidLayers[strCurrLayer] = struct{}{}
		}

		return next
	}

	key := prevLayer[idx : idx+2]
	for _, letter := range trios[key] {
		if nextLayer(prevLayer, append(currLayer, letter), idx+1, invalidLayers, trios) {
			return true
		}
	}

	return false
}

func main() {
	bottom := "AAAA"
	allowed := []string{"AAB", "AAC", "BCD", "BBE", "DEF"}
	fmt.Println(pyramidTransition(bottom, allowed))
}
