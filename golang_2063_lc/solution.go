package main

func countVowels(word string) int64 {
	var res int
	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	for i, letter := range word {
		if _, ok := vowels[letter]; ok {
			res += (i + 1) * (len(word) - i)
		}
	}

	return int64(res)
}

func main() {

}
