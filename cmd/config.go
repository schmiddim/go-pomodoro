package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print Config",
	Long:  `Set config with export POMODORO_API_KEY="MYSUPERSECRETKEY"; export POMODORO_ENDPOINT="http://mykozukai-endpoint.com/api/accomplishments"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("POMODORO_API_KEY: " + viper.GetString("POMODORO_API_KEY"))
		fmt.Println("POMODORO_ENDPOINT: " + viper.GetString("POMODORO_ENDPOINT"))

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
