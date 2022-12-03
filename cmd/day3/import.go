package day3

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "3",
		Short: "Problems for Day 3",
	}

	day.AddCommand(aCmd)
	day.AddCommand(bCmd)

	root.AddCommand(day)
}
