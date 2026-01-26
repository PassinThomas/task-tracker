package cmd

import (
	"errors"

	"task/internal/service"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:	"add",
	Short:	"Command to add to-do list",
	RunE:	func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Wrong arguments")
		}
		err := service.Add(args[0])
		if err != nil {
			return errors.New("Add todo-list failed")
		}
		return nil
	},
}
