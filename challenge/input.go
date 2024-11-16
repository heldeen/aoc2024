package challenge

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Input struct {
	scanner *bufio.Scanner

	lines chan string
}

func FromFileP(p string) (*Input, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("opening input.txt file %w", err)
	}

	return newInputFromReader(f, f), nil
}

func FromLiteral(input string) *Input {
	return newInputFromReader(strings.NewReader(input), nil)
}

func newInputFromReader(r io.Reader, c io.Closer) *Input {
	result := &Input{
		scanner: bufio.NewScanner(r),
		lines:   make(chan string),
	}

	go func() {
		defer func() {
			if c != nil {
				_ = c.Close()
			}
		}()

		for result.scanner.Scan() {
			result.lines <- result.scanner.Text()
		}

		close(result.lines)
	}()

	return result
}

func (c *Input) Lines() <-chan string {
	return c.lines
}
