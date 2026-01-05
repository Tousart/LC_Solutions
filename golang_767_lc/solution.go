package main

import "fmt"

func reorganizeString(s string) string {
	freqs := make([]int, 26)
	var (
		mostFreqIdx int
		maxCount    int
	)

	for _, letter := range s {
		idx := int(letter - 'a')
		freqs[idx]++

		if freqs[idx] > maxCount {
			maxCount = freqs[idx]
			mostFreqIdx = idx
		}
	}

	if maxCount-1 > len(s)-maxCount {
		return ""
	}

	idxsWithFreq := make([]int, 0)
	firstIdx := 0
	for i, freq := range freqs {
		if freq > 0 {
			if i == mostFreqIdx {
				firstIdx = len(idxsWithFreq)
			}
			idxsWithFreq = append(idxsWithFreq, i)
		}
	}
	idxsWithFreq[0], idxsWithFreq[firstIdx] = idxsWithFreq[firstIdx], idxsWithFreq[0]

	res := make([]byte, len(s))

	currIdx := 0
	for i := 0; i < len(s); i += 2 {
		res[i] = byte(idxsWithFreq[currIdx] + 'a')
		freqs[idxsWithFreq[currIdx]]--
		if freqs[idxsWithFreq[currIdx]] == 0 {
			currIdx++
		}
	}

	for i := 1; i < len(s); i += 2 {
		res[i] = byte(idxsWithFreq[currIdx] + 'a')
		freqs[idxsWithFreq[currIdx]]--
		if freqs[idxsWithFreq[currIdx]] == 0 {
			currIdx++
		}
	}

	return string(res)
}

func main() {
	s := "vvvllgo"
	fmt.Println(reorganizeString(s))
}
