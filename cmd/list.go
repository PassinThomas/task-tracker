package cmd

import (
	"fmt"
	// "errors"
	"task/internal/util"

	"github.com/spf13/cobra"
)



var (
	sort bool
	lstCmd = &cobra.Command{
		Use:	"list",
		Short:	"print to-do list",
		RunE:	func(cmd *cobra.Command, args []string) error{
			if sort {
				util.Vlog(util.Verbose, "Sorted actived")
			}
			fmt.Println("Test numero 1")
			return nil
		},
	}
)