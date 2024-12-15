package day3

import (
	"regexp"

	"github.com/heldeen/aoc2024/challenge"
)

func PartB(challenge *challenge.Input) int {

	r := regexp.MustCompile(`(?:do\(\)|don't\(\))|mul\((\d{1,3}),(\d{1,3})\)`)
	total := 0

	do := true
	for l := range challenge.Lines() {
		matches := r.FindAllStringSubmatch(l, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				do = true
			} else if match[0] == "don't()" {
				do = false
			} else if len(match) == 3 && do { // 3 elements: full match, group 1 (x), group 2 (y)
				x := match[1]
				y := match[2]
				total += mustAtoi(x) * mustAtoi(y)
			}
		}
	}

	return total
}
