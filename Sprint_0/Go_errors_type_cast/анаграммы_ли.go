package main

import (
	"fmt"
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	str1Rune := []rune(strings.ToLower(str1))
	str2Rune := []rune(strings.ToLower(str2))

	sort.Sort(sortRunes(str1Rune))
	sort.Sort(sortRunes(str2Rune))

	return strings.EqualFold(string(str1Rune), string(str2Rune))
}

type sortRunes []rune

func (s sortRunes) Len() int           { return len(s) }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	fmt.Println(AreAnagrams("Кабан", "банка"))
}
