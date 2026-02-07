package cmd

import (
	"fmt"

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
		RunE:	func(cmd *cobra.Command, args []string) error {
			err := service.List(option, sorting)
			if err != nil {
				return fmt.Errorf("%w", err)
			}
			return nil
		},
	}
)