package day1

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/heldeen/aoc2024/challenge"
)

func PartA(challenge *challenge.Input) int {

	var listA, listB []int
	for line := range challenge.Lines() {
		s := strings.Fields(line)
		listA = append(listA, mustAtoi(s[0]))
		listB = append(listB, mustAtoi(s[1]))
	}

	slices.Sort(listA)
	slices.Sort(listB)

	totalDist := 0
	for i := 0; i < len(listA); i++ {
		totalDist += int(math.Abs(float64(listA[i] - listB[i])))
	}

	return totalDist
}

func mustAtoi(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}
