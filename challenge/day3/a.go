package day3

import (
	"regexp"
	"strconv"

	"github.com/heldeen/aoc2024/challenge"
)

func PartA(challenge *challenge.Input) int {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	total := 0
	for l := range challenge.Lines() {
		matches := r.FindAllStringSubmatch(l, -1)

		for _, match := range matches {
			if len(match) == 3 { // 3 elements: full match, group 1 (x), group 2 (y)
				x := match[1]
				y := match[2]
				total += mustAtoi(x) * mustAtoi(y)
			}
		}
	}

	return total
}

func mustAtoi(i string) int {
	atoi, err := strconv.Atoi(i)
	if err != nil {
		panic(err)
	}

	return atoi
}
