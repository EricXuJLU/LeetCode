package main

// Day01 1662
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1, s2 := "", ""
	for _, v := range word1 {
		s1 += v
	}
	for _, v := range word2 {
		s2 += v
	}
	return s1 == s2
}
