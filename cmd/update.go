package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:	"update",
	Short:	"Command to update to-do list",
	Run:	func(cmd *cobra.Command, args []string){
		if verbose {
			fmt.Println("mode [verbose] -- RUN")
		}
		fmt.Println("Subcommand update used")
	},
}
