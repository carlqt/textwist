package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"
)

type Word []rune

func (w Word) Len() int           { return len(w) }
func (w Word) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w Word) Less(i, j int) bool { return w[i] < w[j] }

func main() {
	var input string

	str1 := Word("carpenter")

	x := Shuffle(str1)
	fmt.Println(string(x))

	fmt.Print("Enter a word to compare to carpenter: ")
	fmt.Scan(&input)

	if similar(str1, []rune(input)) {
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
	i := sort.Search(len(src), func(i int) bool { return src[i] >= str })

	if i < len(src) && src[i] == str {
		return true
	} else {
		return false
	}
}
func Shuffle(vals Word) Word {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make(Word, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}
