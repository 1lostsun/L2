package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	slice := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	fmt.Println(anagramSearcher(slice))
}

func anagramSearcher(slice []string) map[string][]string {
	mp := make(map[string][]string)
	rep := make(map[string]string)
	for _, v := range slice {
		word := strings.ToLower(v)
		key := runeSorter([]rune(word))

		if _, ok := rep[key]; !ok {
			rep[key] = word
		}

		mp[rep[key]] = append(mp[rep[key]], word)
	}

	for k, v := range mp {
		if len(v) < 2 {
			delete(mp, k)
			continue
		}
		sort.Strings(v)
		mp[k] = v
	}
	return mp
}

func runeSorter(word []rune) string {
	sort.Slice(word, func(i, j int) bool {
		return word[i] < word[j]
	})

	return string(word)
}
