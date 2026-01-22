package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lstCmd = &cobra.Command{
	Use:	"list",
	Short:	"print to-do list",
	Run:	func(cmd *cobra.Command, args []string){
		fmt.Println("subcommand list used")
	},
}