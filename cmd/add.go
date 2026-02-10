package cmd

import (
	"fmt"

	"task/internal/service"
	"task/internal/utils"

	"github.com/spf13/cobra"
)


var addCmd = &cobra.Command{
	Use:	"add",
	Short:	"Command to add to-do list",
	Args: cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE:	func(cmd *cobra.Command, args []string) error {
		err := utils.ParseStr(args[0])
		if err != nil {
			utils.Vlog(fmt.Sprintf("%v", err))
			return fmt.Errorf("Bad format: %w", err)
		}
		err = service.Add(args[0])
		if err != nil {
			return fmt.Errorf("Add todo-list failed: %w", err)
		}
		utils.Vlog("Task %s created" + args[0])
		return nil
	},
}

