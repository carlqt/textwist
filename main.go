package main

import (
	"log"
	"sort"
)

type Word []rune

func (w Word) Len() int           { return len(w) }
func (w Word) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w Word) Less(i, j int) bool { return w[i] < w[j] }

func main() {
	str1 := Word("apples")
	str2 := Word("apple")

	if similar(str1, str2) {
		log.Println("We have the same elements")
	} else {
		log.Println("We are not related :(")
	}
}

func similar(src, word Word) bool {
	srcCpy := make(Word, len(src))
	wordCpy := make(Word, len(word))

	copy(srcCpy, src)
	copy(wordCpy, word)

	sort.Sort(srcCpy)
	sort.Sort(wordCpy)

	for _, w := range wordCpy {
		if !contains(srcCpy, w) {
			return false
		}
	}

	return true
}

func contains(src Word, str rune) bool {
	for _, a := range src {
		if a == str {
			return true
		}
	}

	return false
}
