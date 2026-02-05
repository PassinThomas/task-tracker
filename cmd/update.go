package cmd

import (
	"fmt"

	"task/internal/utils"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:	"update",
	Short:	"Command to update to-do list",
	Run:	func(cmd *cobra.Command, args []string){
		if utils.Verbose {
			fmt.Println("mode [verbose] -- RUN")
		}
		fmt.Println("Subcommand update used")
	},
}
