package day1

import (
	"strings"

	"github.com/heldeen/aoc2024/challenge"
)

func PartB(challenge *challenge.Input) int {
	var listA []int
	setB := make(set)
	for line := range challenge.Lines() {
		s := strings.Fields(line)
		listA = append(listA, mustAtoi(s[0]))
		setB.Add(mustAtoi(s[1]))
	}

	totalSimilarity := 0
	for _, a := range listA {
		totalSimilarity += a * setB.Get(a)
	}

	return totalSimilarity
}

type set map[int]int

func (s set) Len() int {
	return len(s)
}

func (s set) Add(i int) {
	v, ok := s[i]
	if ok {
		s[i] = v + 1
		return
	}
	s[i] = 1
}

func (s set) Get(i int) int {
	return s[i]
}

func (s set) Contains(i int) bool {
	_, ok := s[i]
	return ok
}
