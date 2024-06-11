package cmd

import (
	"fmt"

	"github.com/silvano-paulino/internal/application"
	"github.com/silvano-paulino/internal/domain"
	"github.com/spf13/cobra"
)

var userId string

var fetchUserCmd = &cobra.Command{
	Use:   "fetchuser",
	Short: "Fetch a user by ID",
	Long:  `Fetch a user by their ID from the mock database.`,
	Run: func(cmd *cobra.Command, args []string) {
		mockRepo := &MockDatabase{
			MockUserData: map[string]*domain.User{
				"1": {Id: "1", Name: "Silvano Paulino"},
			},
		}

		user, err := application.FetchUser(mockRepo, userId)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("User ID: %s, Name: %s\n", user.Id, user.Name)
	},
}

func init() {
	rootCmd.AddCommand(fetchUserCmd)
	fetchUserCmd.Flags().StringVarP(&userId, "id", "i", "", "ID of the user to fetch")
}
