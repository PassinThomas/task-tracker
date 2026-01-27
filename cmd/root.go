package cmd

import (
	"task/internal/util"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:	"task",
		Short:	"task tool",
		Long:	"software to track to-do lists",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&util.Verbose,
		"verbose",
		"v",
		false,
		"verbose mode",
	)
	lstCmd.Flags().BoolVarP(&sort, "sort", "s", false, `sort todo-list ("name", "date", "done")`,)
	rootCmd.AddCommand(addCmd, deleteCmd, updateCmd, lstCmd)
}

func Execute() {
	rootCmd.SilenceErrors = true
	cobra.CheckErr(rootCmd.Execute())
}


