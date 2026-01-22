package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:	"add",
	Short:	"Command to add to-do list",
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Println("Subcommand add used")
	},
}

