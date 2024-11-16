package cmd

import (
	"strconv"

	"github.com/heldeen/aoc2024/gen"
	"github.com/spf13/cobra"
)

func init() {
	genCmd := &cobra.Command{
		Use:       "gen <day>",
		Short:     "generate the boilerplate for the specified day's AoC solution",
		Example:   "gen 11",
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		ValidArgs: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25"},
		RunE: func(cmd *cobra.Command, args []string) error {
			day, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			return gen.GenerateDay(day)
		},
	}

	rootCmd.AddCommand(genCmd)
}
