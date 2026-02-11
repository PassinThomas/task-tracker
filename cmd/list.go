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
	option	string
	sorting	string
	lstCmd = &cobra.Command{
		Use:	"list",
		Short:	"print to-do list",
		Long:  "Print to-do list. Use --status=done or --status=not-done to filter.",
		SilenceUsage: true,
		RunE:	func(cmd *cobra.Command, args []string) error {
			utils.Debug("Run list process...",
				"cmd", "list",
				"diplay_option", option,
				"sorting", sorting,
			)
			tasks, err := service.List(option, sorting)
			if err != nil {
				return err
			}
			done := "not-done"

			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"ID", "Title", "Status", "Created At"})
			for _, task := range tasks {
				if task.Done {
					done = "done"
				}
				table.Append([]string{
					fmt.Sprintf("%v", task.ID),
					task.Title,
					done,
					task.CreatedAt.Format("2006-01-02 15:04"),
				})
			}
			table.Render()
			return nil
		},
	}
)