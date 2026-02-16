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
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
            // Ici, utils.DebugVar contient la vraie valeur passée par l'utilisateur
            utils.InitLogger(utils.DebugVar)
        },
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&utils.DebugVar,
		"debug",
		"d",
		false,
		"debug mode enabled",
	)

	lstCmd.Flags().StringVarP(&status, "status", "s", "", `display-list ("not-done", "done")`,)
	lstCmd.Flags().StringVarP(&sorting, "sort", "", "", "specify the field to sort tasks by; valid options are 'title', 'created' (newest first), 'updated', or 'status'")
	lstCmd.Flags().StringVarP(&order, "order", "", "",  "set the sort order for tasks; options are 'asc' for ascending or 'desc' for descending")
	updateCmd.Flags().StringVarP(&newTitle, "title", "", "", "update title of task")
	updateCmd.Flags().BoolVarP(&markDone, "done", "", false, "mark done task finished")
	rootCmd.AddCommand(addCmd, deleteCmd, updateCmd, lstCmd)
}

func Execute() {
	rootCmd.SilenceErrors = true
	utils.InitLogger(utils.DebugVar)
	cobra.CheckErr(rootCmd.Execute())
}


