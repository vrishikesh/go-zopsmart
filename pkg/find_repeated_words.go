package pkg

import (
	"sort"
	"strings"
)

func FindRepeatedWords(input string) []string {
	m := make(map[string]int)
	for _, w := range strings.Fields(input) {
		m[w] += 1
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys[:3]
}
