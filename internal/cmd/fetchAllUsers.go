package cmd

import (
	"fmt"

	"github.com/silvano-paulino/internal/application"
	"github.com/silvano-paulino/internal/domain"
	"github.com/spf13/cobra"
)

var fetchAllUsersCmd = &cobra.Command{
	Use:   "fetchallusers",
	Short: "Fetch all users",
	Long:  `Fetch all users from the mock database.`,
	Run: func(cmd *cobra.Command, args []string) {
		mockRepoAll := &MockDatabase{
			MockAllUserData: []domain.User{
				{Id: "1", Name: "John Doe"},
				{Id: "2", Name: "Jane Smith"},
				{Id: "3", Name: "Alice Johnson"},
			},
		}

		users, err := application.FetchAllUsers(mockRepoAll)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		for _, user := range users {
			fmt.Printf("User ID: %s, Name: %s\n", user.Id, user.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchAllUsersCmd)
}
