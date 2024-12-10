package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2024/challenge"
	"github.com/heldeen/aoc2024/challenge/day2"
)

func init() {
	const inputFlag = "input"
	const inputFlagShort = "i"
	const inputFlagUsage = "path of the input file to use"

	var inputFlagValue string

	day := &cobra.Command{
		Use:   "2",
		Short: "Problems for Day 2",
	}

	a := &cobra.Command{
		Use:   "a",
		Short: "Day 2, Problem A",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			input, err := challenge.FromFileP(flag.Value.String())
			cobra.CheckErr(err)
			fmt.Printf("Day 2, Part A - Answer: %v\n", day2.PartA(input))
		},
	}

	a.Flags().StringVarP(&inputFlagValue, inputFlag, inputFlagShort, "./challenge/day2/input.txt", inputFlagUsage)

	day.AddCommand(a)

	b := &cobra.Command{
		Use:   "b",
		Short: "Day 2, Part B",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			input, err := challenge.FromFileP(flag.Value.String())
			cobra.CheckErr(err)
			fmt.Printf("Day 2, Part B - Answer: %v\n", day2.PartB(input))
		},
	}

	b.Flags().StringVarP(&inputFlagValue, inputFlag, inputFlagShort, "./challenge/day2/input.txt", inputFlagUsage)

	day.AddCommand(b)

	rootCmd.AddCommand(day)
}
