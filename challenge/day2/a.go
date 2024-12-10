package day2

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/heldeen/aoc2024/challenge"
)

func PartA(challenge *challenge.Input) int {
	saveLevels := 0

	for l := range challenge.Lines() {
		levelsStr := strings.Fields(l)
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
		if inc || dec {
			saveLevels++
			log.Println(levelsStr)
		}
	}

	return saveLevels
}

func mustAtoi(i string) int {
	atoi, err := strconv.Atoi(i)
	if err != nil {
		panic(err)
	}

	return atoi
}
