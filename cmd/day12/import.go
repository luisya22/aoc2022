package day12

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "12",
		Short: "Problems for Day 12",
	}

	day.AddCommand(aCmd)
	day.AddCommand(bCmd)

	root.AddCommand(day)
}
