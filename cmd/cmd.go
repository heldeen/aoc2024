package cmd

import (
	"fmt"
	"time"

	"github.com/pkg/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type prof interface {
	Stop()
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var (
	start    time.Time
	profiler prof

	rootCmd = &cobra.Command{
		Use:     "aoc2024",
		Short:   "Advent of Code 2024 Solutions",
		Long:    "Golang implementations for the 2024 Advent of Code problems",
		Example: "./aoc2024 1 a -i ./challenge/day1/input.txt.txt",
		Args:    cobra.ExactArgs(1),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if viper.GetBool("profile") {
				profiler = profile.Start()
			}

			start = time.Now()
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			if profiler != nil {
				profiler.Stop()
			}

			fmt.Println("Took", time.Since(start))
		},
	}
)

func init() {

	flags := rootCmd.PersistentFlags()

	flags.Bool("profile", false, "Profile implementation performance")

	_ = viper.BindPFlags(flags)
}
