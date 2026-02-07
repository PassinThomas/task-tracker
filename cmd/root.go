package cmd

import (
	"task/internal/utils"

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
		&utils.Verbose,
		"verbose",
		"v",
		false,
		"verbose mode",
	)
	lstCmd.Flags().StringVarP(&option, "status", "s", "", `display-list ("not-done", "done")`,)
	lstCmd.Flags().StringVarP(&sorting, "sort", "", "", `sort-task by ("title", "date(by Newest)", "status")`)
	updateCmd.Flags().StringVarP(&newTitle, "title", "", "", "update title of task")
	updateCmd.Flags().BoolVarP(&markDone, "done", "", false, "mark done task finished")
	rootCmd.AddCommand(addCmd, deleteCmd, updateCmd, lstCmd)
}

func Execute() {
	rootCmd.SilenceErrors = true
	cobra.CheckErr(rootCmd.Execute())
}


