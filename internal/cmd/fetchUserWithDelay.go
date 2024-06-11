package cmd

import (
	"fmt"
	"time"

	"github.com/silvano-paulino/internal/application"
	"github.com/silvano-paulino/internal/domain"
	"github.com/spf13/cobra"
)

var userIdWithDelay string

var fetchUserWithDelayCmd = &cobra.Command{
	Use:   "fetchuserwithdelay",
	Short: "Fetch a user by ID with a simulated delay",
	Long:  `Fetch a user by their ID from the mock database, simulating a delay in the response.`,
	Run: func(cmd *cobra.Command, args []string) {
		mockRepo := &MockDatabase{
			MockUserData: map[string]*domain.User{
				"4": {Id: "4", Name: "Silvano Paulino"},
			},
		}

		fmt.Println("Fetching user, please wait...")
		start := time.Now()
		user, err := application.FetchUserWithDelay(mockRepo, userIdWithDelay)
		duration := time.Since(start)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Fetched user in %v\n", duration)
		fmt.Printf("User ID: %s, Name: %s\n", user.Id, user.Name)
	},
}

func init() {
	rootCmd.AddCommand(fetchUserWithDelayCmd)
	fetchUserWithDelayCmd.Flags().StringVarP(&userIdWithDelay, "id", "i", "", "ID of the user to fetch")
}
