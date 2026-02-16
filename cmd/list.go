package cmd

import (
	"fmt"
	"os"

	"task/internal/service"
	"task/internal/utils"

	"github.com/spf13/cobra"
	"github.com/olekukonko/tablewriter"

)

var (
	status	string
	sorting	string
	order	string
	lstCmd = &cobra.Command{
		Use:	"list",
		Short:	"print to-do list",
		Long: `Print your to-do list with optional filtering and sorting.

You can filter tasks with the following flags:
  --done          : show only tasks marked as done
  --not-done      : show only tasks not yet done
  --updated       : show only tasks that have been updated
  --not-updated   : show only tasks that have never been updated

You can sort tasks with:
  --sort <field>  : field to sort by ('title', 'created', 'updated', 'status')
  --order <asc|desc> : sort order, ascending or descending

Examples:
  task list --done --sort created --order desc
  task list --not-done --sort title --order asc`,
		SilenceUsage: true,
		Args:	cobra.ExactArgs(0),
		RunE:	func(cmd *cobra.Command, args []string) error {
			utils.Debug("Run list process...",
				"cmd", "list",
				"diplay_option", status,
				"sorting", sorting,
				"order", order,
			)
			tasks, err := service.List()
			if err != nil {
				return err
			}
			if sorting != "" {
				utils.OptSort[sorting] = order
				tasks, err = service.SortList(tasks, utils.OptSort)
				if err != nil {
					return err
				}
			}
			done := "not-done"

			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"ID", "Title", "Status", "Created At", "Updated At"})
			for _, task := range tasks {
				if task.Done {
					done = "done"
				}
				table.Append([]string{
					fmt.Sprintf("%v", task.ID),
					task.Title,
					done,
					task.CreatedAt.Format("2006-01-02 15:04"),
					utils.UpadtedFormat(task),
				})
			}
			table.Render()
			return nil
		},
	}
)