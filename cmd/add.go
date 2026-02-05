package cmd

import (
	"fmt"
	"errors"

	"task/internal/service"
	"task/internal/utils"

	"github.com/spf13/cobra"
)


var addCmd = &cobra.Command{
	Use:	"add",
	Short:	"Command to add to-do list",
	RunE:	func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Wrong arguments")
		}
		err := utils.ParseStr(args[0])
		if err != nil {
			utils.Vlog(utils.Verbose, fmt.Sprintf("%v", err))
			return fmt.Errorf("Incorrect format: %w", err)
		}
		err = service.Add(args[0])
		if err != nil {
			return fmt.Errorf("Add todo-list failed: %w", err)
		}
		utils.Vlog(utils.Verbose, "Task %s created" + args[0])
		return nil
	},
}

