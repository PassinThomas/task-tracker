package cmd

import (
	"github.com/PassinThomas/task-tracker/models"
	"github.com/PassinThomas/task-tracker/internal/utils"

	"github.com/spf13/cobra"
)

var (
	sorting	string
	order	string
	lstCmd = &cobra.Command{
		Use:	"list",
		Short:	"print to-do list",
		Long: `Print your to-do list with optional filtering and sorting.

You can filter tasks with the following flags:
  --done          : show only tasks marked as done
  --undone		  : show only tasks not yet done
  --updated       : show only tasks that have been updated
  --not-updated   : show only tasks that have never been updated

You can sort tasks with:
  --sort <field>  : field to sort by ('title', 'created', 'updated')
  --order <asc|desc> : sort order, ascending or descending

Examples:
  task list --done --sort created --order desc
  task list --not-done --sort title --order asc`,
		SilenceUsage: true,
		Args:	cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {	
			utils.Debug("Run list process...",
				"cmd", "list",
				"display_option", filter,
				"sorting", sorting,
				"order", order,
			)
		
			opts := models.ListOptions{
				Filter: *filter,
				Sort:   sorting,
				Order:  order,
			}
		
			tasks, err := taskService.ListWithOptions(opts)
			if err != nil {
				return err
			}
			
			utils.Debug("todo-list", tasks)
			utils.RenderTasks(tasks)
			
			return nil
		},
	}
)