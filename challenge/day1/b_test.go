package day1

import (
	"testing"

	"github.com/heldeen/aoc2024/challenge"
)

func TestB(t *testing.T) {
	want := 31

	input := challenge.FromLiteral(sample)

	result := PartB(input)

	if result != want {
		t.Errorf("Day[1] Part[B] - wanted [%d] but got [%d]", want, result)
	}
}
