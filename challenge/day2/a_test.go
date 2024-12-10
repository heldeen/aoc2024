package day2

import (
	"testing"

	"github.com/heldeen/aoc2024/challenge"
)

const sample = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestA(t *testing.T) {
	want := 2

	input := challenge.FromLiteral(sample)

	result := PartA(input)

	if result != want {
		t.Errorf("Day[2] Part[A] - wanted [%d] but got [%d]", want, result)
	}
}
