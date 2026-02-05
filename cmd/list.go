package cmd

import (
	// "fmt"
	"errors"
	// "task/internal/utils"
	"task/internal/service"

	"github.com/spf13/cobra"
)



var (
	option	string
	sorting	string
	lstCmd = &cobra.Command{
		Use:	"list",
		Short:	"print to-do list",
		Long:  "Print to-do list. Use --status=done or --status=not-done to filter.",
		// Long:  `Print to-do list. Use --sort=("title", "date", "status").`,
		RunE:	func(cmd *cobra.Command, args []string) error{
			if len(args) != 0 {
				return errors.New("Wrong number of arguments")
			}
			service.List(option, sorting)
			return nil
		},
	}
)