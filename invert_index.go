package main

import (
	"strings"
)

type InvertedIndex map[string][]string

func BuildInvertedIndex(documents []string, id_cnt string) InvertedIndex {
	index := make(InvertedIndex)

	for _, doc := range documents {
		docID := id_cnt
		words := strings.Fields(doc)

		for _, word := range words {
			word = strings.ToLower(word)

			if _, ok := index[word]; !ok {
				word_last_index, err := list_end(word)
				if err != nil {
					index[word] = []string{docID}
				} else if docID != word_last_index {
					index[word] = []string{docID}
				}

			}
		}
	}

	return index
}

func contains(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
