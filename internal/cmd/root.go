package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "usercli",
	Short: "CLI for user operations",
	Long:  `A Command Line Interface for performing operations on users.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func init() {
	// Register commands here
	rootCmd.AddCommand(fetchUserCmd)
	rootCmd.AddCommand(fetchAllUsersCmd)
	rootCmd.AddCommand(fetchUserWithDelayCmd)
}
