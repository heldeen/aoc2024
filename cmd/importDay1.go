package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2024/challenge"
	"github.com/heldeen/aoc2024/challenge/day1"
)

func init() {
	const inputFlag = "input"
	const inputFlagShort = "i"
	const inputFlagUsage = "path of the input file to use"

	var inputFlagValue string

	day := &cobra.Command{
		Use:   "1",
		Short: "Problems for Day 1",
	}

	a := &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			input, err := challenge.FromFileP(flag.Value.String())
			cobra.CheckErr(err)
			fmt.Printf("Day 1, Part A - Answer: %v\n", day1.PartA(input))
		},
	}

	a.Flags().StringVarP(&inputFlagValue, inputFlag, inputFlagShort, "./challenge/day1/input.txt", inputFlagUsage)

	day.AddCommand(a)

	b := &cobra.Command{
		Use:   "b",
		Short: "Day 1, Part B",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			input, err := challenge.FromFileP(flag.Value.String())
			cobra.CheckErr(err)
			fmt.Printf("Day 1, Part B - Answer: %v\n", day1.PartB(input))
		},
	}

	b.Flags().StringVarP(&inputFlagValue, inputFlag, inputFlagShort, "./challenge/day1/input.txt", inputFlagUsage)

	day.AddCommand(b)

	rootCmd.AddCommand(day)
}
