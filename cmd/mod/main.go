package main

import (
	"fmt"

	"github.com/itsmikoj/mod-go/cmd/mod/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mod",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello!")
	},
}

func main() {
	getCmd := commands.NewGetCommand()
	removeCmd := commands.NewRemoveCommand()
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.Execute()
}
