package day1

import (
	"testing"

	"github.com/heldeen/aoc2024/challenge"
)

const sample = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestA(t *testing.T) {
	want := 11

	input := challenge.FromLiteral(sample)

	result := PartA(input)

	if result != want {
		t.Errorf("Day[1] Part[A] - wanted [%d] but got [%d]", want, result)
	}
}
