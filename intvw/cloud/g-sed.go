package intview

import (
	s "strings"
)

func EchoPipeSed(word1 string, word2 string) string {
	word3 := ""

	for _, c := range word1 {
		letter := string(c)
		if s.Contains(word2, letter) {
			word2 = s.Replace(word2, letter, "")
		} else {
			if !s.Contains(word3, letter) {
				word3 += letter
			}
		}
	}
	word3 += word2
	return word3
}
