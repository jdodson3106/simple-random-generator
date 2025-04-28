package srg

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	runs := 100_000
	resultStore := make([]string, runs)
	genreateData(&resultStore)

	first := AlphaNums[0]
	last := AlphaNums[len(AlphaNums)-1]
	hasFirst := false
	hasLast := false
	for _, result := range resultStore {
		if hasFirst && hasLast {
			break
		}

		for _, b := range result {
			if b == rune(first) {
				hasFirst = true
				continue
			} else if b == rune(last) {
				hasLast = true
				continue
			}
		}
	}

	assert.True(t, hasFirst)
	assert.True(t, hasLast)
}

func genreateData(resultStore *[]string) {
	var wg sync.WaitGroup

	results := make(chan string)

	for i := 0; i < len(*resultStore); i++ {
		wg.Add(1)
		go generator(&wg, results)
	}

	for res := 0; res < len(*resultStore); res++ {
		(*resultStore)[res] = <-results
	}
}

func generator(wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	val := GenerateAlphaNumberic(21)
	results <- val
}
