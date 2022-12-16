package cmd

// This code generated by go generate.
// DO NOT EDIT BY HAND!

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/luisya22/aoc2022/cmd/day1"
	"github.com/luisya22/aoc2022/cmd/day10"
	"github.com/luisya22/aoc2022/cmd/day11"
	"github.com/luisya22/aoc2022/cmd/day12"
	"github.com/luisya22/aoc2022/cmd/day13"
	"github.com/luisya22/aoc2022/cmd/day14"
	"github.com/luisya22/aoc2022/cmd/day2"
	"github.com/luisya22/aoc2022/cmd/day3"
	"github.com/luisya22/aoc2022/cmd/day4"
	"github.com/luisya22/aoc2022/cmd/day5"
	"github.com/luisya22/aoc2022/cmd/day6"
	"github.com/luisya22/aoc2022/cmd/day7"
	"github.com/luisya22/aoc2022/cmd/day8"
	"github.com/luisya22/aoc2022/cmd/day9"
)

func addDays(root *cobra.Command) {

	day1.AddCommandsTo(root)
	day10.AddCommandsTo(root)
	day11.AddCommandsTo(root)
	day12.AddCommandsTo(root)
	day13.AddCommandsTo(root)
	day14.AddCommandsTo(root)
	day2.AddCommandsTo(root)
	day3.AddCommandsTo(root)
	day4.AddCommandsTo(root)
	day5.AddCommandsTo(root)
	day6.AddCommandsTo(root)
	day7.AddCommandsTo(root)
	day8.AddCommandsTo(root)
	day9.AddCommandsTo(root)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc2022",
	Short: "Advent of Code 2022 Solutions",
	Long:  `Golang implementations for the 2022 Advent of Code problems.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aoc2022.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	addDays(rootCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
