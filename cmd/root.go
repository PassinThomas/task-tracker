package cmd

import (
	"github.com/PassinThomas/task-tracker/internal/utils"
	"github.com/PassinThomas/task-tracker/internal/service"
	"github.com/PassinThomas/task-tracker/internal/store"
	"github.com/PassinThomas/task-tracker/models"

	"github.com/spf13/cobra"
)

var (
	flg = &models.FlgUpdate{}
	filter = &models.FilterOptions{}
	taskService *service.TaskService
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

	lstCmd.Flags().BoolVarP(&filter.Done, "done", "", false, `display-list "done"`,)
	lstCmd.Flags().BoolVarP(&filter.Undone, "undone", "", false, `display-list "undone"`,)
	lstCmd.MarkFlagsMutuallyExclusive("done", "undone")
	
	lstCmd.Flags().BoolVarP(&filter.Updated, "updated", "", false, `display-list "updated"`,)
	lstCmd.Flags().BoolVarP(&filter.NotUpdated, "not-updated", "", false, `display-list "not updated"`)
	lstCmd.MarkFlagsMutuallyExclusive("updated", "not-updated")
	
	lstCmd.Flags().StringVarP(&sorting, "sort", "", "", "specify the field to sort tasks by; valid options are 'title', 'created' (newest first), 'updated', or 'status'")
	lstCmd.Flags().StringVarP(&order, "order", "", "",  "set the sort order for tasks; options are 'asc' for ascending or 'desc' for descending")
	
	updateCmd.Flags().StringVarP(&flg.NewTitle, "title", "", "", "update title of task")
	updateCmd.Flags().BoolVarP(&flg.Done, "complete", "c", false, "mark done task finished")
	updateCmd.Flags().BoolVarP(&flg.NotDone, "incomplete", "i", false, "mark undone task not finished")
	updateCmd.MarkFlagsMutuallyExclusive("complete", "incomplete")

	rootCmd.AddCommand(addCmd, deleteCmd, updateCmd, lstCmd)
}

func Execute() {
	jsonStore := store.NewJsonStore()
	taskService = service.NewTaskService(jsonStore)

	rootCmd.SilenceErrors = true
	utils.InitLogger(utils.DebugVar)
	cobra.CheckErr(rootCmd.Execute())
}


