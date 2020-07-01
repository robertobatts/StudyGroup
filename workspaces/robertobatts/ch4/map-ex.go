package main

import (
	"bufio"
	"fmt"
	"io"
)

func scanWords(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	words := make(map[string]int)
	for scanner.Scan() {
		words[scanner.Text()]++
	}
	for word, count := range words {
		fmt.Printf("%v\t%v\n", count, word)
	}
}
