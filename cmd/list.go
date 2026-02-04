package cmd

import (
	// "fmt"
	// "errors"
	// "task/internal/util"
	"task/internal/service"

	"github.com/spf13/cobra"
)



var (
	option string
	lstCmd = &cobra.Command{
		Use:	"list",
		Short:	"print to-do list",
		 Long:  "Print to-do list. Use --status=done or --status=not-done to filter.",
		RunE:	func(cmd *cobra.Command, args []string) error{
			service.List(&option)
			return nil
		},
	}
)