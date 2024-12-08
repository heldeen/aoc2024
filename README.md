# Advent of Code 2024

[![License](https://img.shields.io/badge/license-MIT-brightgreen)](./LICENSE)

Solutions for the 2024 Advent of Code

## Building

This project makes use of Go 1.23.3.

```bash
git clone git@github.com:heldeen/aoc2024.git
go test ./...
go build
```

## Running the Solutions

To run a solution, use the problem name followed by the path to an input file.

For example, to run problem `2 a`:

```bash
$ go run ./main.go 2 a -i ./challenge/day2/input.txt
Day 2, Part A - Answer: 9633
Took 999.4µs
```

NOTE: Input file is optional and, if omitted, it's location will be inferred by the day parameter to be a file named `input.txt` located in the `challenge/dayX/` folder.

## Adding New Solutions

A generator program is included 

```bash
$ go run ./main.go gen 7
Took 501.833991ms
```

This executes the logic in `gen/problem.go` that makes templates for each day. For
example, `go run main.go gen 9` will generate the following files:

* `cmd/importDay9.go`: A "glue" file combining commands for both of the day's problems to simplify wiring up subcommands
* `challenge/day9/a.go`: The main problem implementation stubbed out for Part A: `func A(*challenge.Input) int`
* `challenge/day9/a_test.go`: A basic test template for Part A
    * This contains a `const sample` that is meant to get the sample input from the AoC website problem description. It uses that to feed the solution input for testing.
* `challenge/day9/b.go`: The main problem implementation stubbed out for Part B: `func B(*challenge.Input) int`
* `challenge/day9/b_test.go`: A basic test template
    * This references the `sample` constant in the same day's `a_test.go` for testing the Part B solution. 
* `challenge/day9/input.txt`: The challenge input downloaded from the AoC website. Grab your https://www.adventofcode.com `session` cookie value and store it in `~/.tokenfile` to enable this.
* `.idea/runtimeConfigurations/TestA_in_aoc2022_challenge_day9.xml`: The JetBrains Goland configuration for running the tests in `a_test.go`.
* `.idea/runtimeConfigurations/TestB_in_aoc2022_challenge_day9.xml`: The JetBrains Goland configuration for running the tests in `b_test.go`.
* `.idea/runtimeConfigurations/Run_aoc2022_challenge_day9_partA.xml.xml`: The JetBrains Goland configuration for running the solution in `a.go` against the input.
* `.idea/runtimeConfigurations/Run_aoc2022_challenge_day9_partB.xml.xml`: The JetBrains Goland configuration for running the solution in `b.go` against the input.

## License

These solutions are licensed under the MIT License.

See [LICENSE](./LICENSE) for details.