// map sorting by vypr
package main

import "sort"

type smap struct {
	Key   string
	Value int
}

type smapList []smap

func (s smapList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s smapList) Len() int {
	return len(s)
}

func (s smapList) Less(i, j int) bool {
	a, b := s[i].Value, s[j].Value

	if a != b {
		return a > b
	} else {
		return s[j].Value > s[i].Value
	}
}

func sortMapByValue(m map[string]int) smapList {
	s := make(smapList, len(m))
	i := 0

	for k, v := range m {
		s[i] = smap{k, v}
		i++
	}

	sort.Sort(s)
	return s
}
