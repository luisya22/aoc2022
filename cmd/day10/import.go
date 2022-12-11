package day10

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "10",
		Short: "Problems for Day 10",
	}

	day.AddCommand(aCmd)
	day.AddCommand(bCmd)

	root.AddCommand(day)
}
