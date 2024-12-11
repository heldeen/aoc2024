package day3

import (
	"testing"

	"github.com/heldeen/aoc2024/challenge"
)

const sample = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func TestA(t *testing.T) {
	want := 161

	input := challenge.FromLiteral(sample)

	result := PartA(input)

	if result != want {
		t.Errorf("Day[3] Part[A] - wanted [%d] but got [%d]", want, result)
	}
}
