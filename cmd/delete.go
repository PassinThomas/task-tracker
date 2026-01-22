package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:	"delete",
	Short:	"Command to delete to-do list",
	Run:	func(cmd *cobra.Command, args []string){
		fmt.Println("Subcommand delete used")
	},
}
