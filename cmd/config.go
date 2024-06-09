package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print Config",
	Long:  `Set config with Environment Variables POMODORO_API_KEY and POMODORO_ENDPOINT`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("POMODORO_API_KEY: " + viper.GetString("POMODORO_API_KEY"))
		fmt.Println("POMODORO_ENDPOINT: " + viper.GetString("POMODORO_ENDPOINT"))

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
