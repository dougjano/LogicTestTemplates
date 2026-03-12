package main

func CheckPhrases(phrase1 []string, phrase2 []string) bool {

	charMap := map[string]int{}

	for _, word := range phrase1 {
		for _, letter := range word {
			charMap[string(letter)] += 1
		}
	}

	for _, word := range phrase2 {
		for _, letter := range word {
			if charMap[string(letter)] > 0 {
				charMap[string(letter)] -= 1
			} else {
				return false
			}
		}
	}

	return true
}
