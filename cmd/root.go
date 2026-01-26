package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	verbose bool
	rootCmd = &cobra.Command{
		Use:	"task",
		Short:	"task tool",
		Long:	"software to track to-do lists",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Main command")
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"verbose mode",
	)
	rootCmd.AddCommand(addCmd, deleteCmd, updateCmd, lstCmd)
}

func Execute() {
	rootCmd.SilenceErrors = true
	cobra.CheckErr(rootCmd.Execute())
}


