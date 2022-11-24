package main

import (
	"bytes"
	"sort"
	"strings"
)

func frequencySort1(s string) string {
	cache := make(map[rune]int, 0)

	for _, v := range s {
		cache[v]++
	}
	slice := make([]kv, 0)
	for k, v := range cache {
		slice = append(slice, kv{k, v})
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].v-slice[j].v > 0
	})

	var result bytes.Buffer

	for _, v := range slice {
		result.WriteString(strings.Repeat(string(v.k), v.v))
	}
	return result.String()

}

type kv struct {
	k rune
	v int
}
