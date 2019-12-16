/*
Karhutin is a simple Markov chain text generator.

It uses an online corpus that contains content from the Karhu Helsinki blog.
*/
package main

import (
	"bufio"
	"fmt"
	"github.com/eminano/markov"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// loadCorpus loads the online corpus from karhutin.surge.sh.
// The corpus is read into a string array. Each array item represents a line
// in the corpus.
func loadCorpus() ([]string, error) {
	res, err := http.Get("https://karhutin.surge.sh/corpus.txt")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var lines []string
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// shuffle shuffles a string array.
// This function is used to shuffle the corpus.
func shuffle(src []string) []string {
	final := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}

	return final
}

func main() {
	chain, _ := markov.NewNGramChain(3)
	lines, err := loadCorpus()

	if err != nil {
		fmt.Print(err)
	}

	shuffled := shuffle(lines)

	chain.ProcessText(strings.NewReader(strings.Join(shuffled[:], "\n")))
	generated := chain.GenerateRandomText(50)

	fmt.Printf("%s\n", generated)
}
