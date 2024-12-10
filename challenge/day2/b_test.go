package day2

import (
	"testing"

	"github.com/heldeen/aoc2024/challenge"
)

func TestB(t *testing.T) {
	want := 42

	input := challenge.FromLiteral(sample)

	result := PartB(input)

	if result != want {
		t.Errorf("Day[2] Part[B] - wanted [%d] but got [%d]", want, result)
	}
}
