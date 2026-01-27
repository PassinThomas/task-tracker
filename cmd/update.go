package cmd

import (
	"fmt"

	"task/internal/util"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:	"update",
	Short:	"Command to update to-do list",
	Run:	func(cmd *cobra.Command, args []string){
		if util.Verbose {
			fmt.Println("mode [verbose] -- RUN")
		}
		fmt.Println("Subcommand update used")
	},
}
