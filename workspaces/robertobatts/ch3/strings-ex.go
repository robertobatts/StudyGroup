package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(anagrams2("abcdeeefgg", "eefgebgacd"))
}

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	dot := strings.LastIndex(s, ".")
	if dot == -1 {
		dot = len(s)
	}
	return s[slash+1 : dot]
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	dot := strings.LastIndex(s, ".")
	if dot == -1 {
		return comma(s[:n-3]) + "," + s[n-3:]
	} else {
		return comma(s[:dot]) + s[dot:]
	}
}

func anagrams1(s1, s2 string) bool {
	var r1 aRune = toRuneSlice(s1)
	var r2 aRune = toRuneSlice(s2)

	sort.Sort(r1)
	sort.Sort(r2)
	return string(r1) == string(r2)
}

type aRune []rune

func (r aRune) Len() int           { return len(r) }
func (r aRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r aRune) Less(i, j int) bool { return r[i] < r[j] }

func toRuneSlice(s string) []rune {
	var r []rune
	for _, runeVal := range s {
		r = append(r, runeVal)
	}
	return r
}

func anagrams2(s1, s2 string) bool {
	sMap1 := toMap(s1)
	sMap2 := toMap(s2)

	for r, count := range sMap1 {
		if sMap2[r] != count {
			return false
		}
	}

	return true
}

func toMap(s string) map[rune]int {
	sMap := make(map[rune]int)
	for _, r := range s {
		sMap[r]++
	}
	return sMap
}
