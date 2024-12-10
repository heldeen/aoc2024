package day2

import (
	"math"
	"strings"

	"github.com/heldeen/aoc2024/challenge"
)

func PartB(challenge *challenge.Input) int {
	safeLevels := 0

	for l := range challenge.Lines() {
		levelsStr := strings.Fields(l)
		if calcSafe(levelsStr) {
			safeLevels++
			continue
		}
		for i := range levelsStr {
			tempLevelsStr := make([]string, 0, len(levelsStr)-1)
			tempLevelsStr = append(tempLevelsStr, levelsStr[:i]...)
			tempLevelsStr = append(tempLevelsStr, levelsStr[i+1:]...)
			if calcSafe(tempLevelsStr) {
				safeLevels++
				break
			}
		}
	}
	return safeLevels
}

func calcSafe(levelsStr []string) bool {
	var prev int
	inc, dec := true, true
	for i, levelStr := range levelsStr {

		cur := mustAtoi(levelStr)

		if i > 0 {
			if cur > prev {
				dec = false
			} else if cur < prev {
				inc = false
			}
			if math.Abs(float64(cur-prev)) < 1 || math.Abs(float64(cur-prev)) > 3 {
				dec = false
				inc = false
				break
			}
		}
		prev = cur
	}
	return inc || dec
}
