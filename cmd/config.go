package cmd

import (
	"fmt"
	"github.com/schmiddim/go-pomodoro/libs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print Auth + Endpoint Config",
	Long:  `Set config with export POMODORO_API_KEY="MYSUPERSECRETKEY"; export POMODORO_ENDPOINT="http://mykozukai-endpoint.com/api/accomplishments"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Configuration:")
		fmt.Println("POMODORO_API_KEY: " + viper.GetString("POMODORO_API_KEY"))
		fmt.Println("POMODORO_ENDPOINT: " + viper.GetString("POMODORO_ENDPOINT"))
		r := libs.NewRest()
		result := r.ConfigTest()
		if result != nil {
			fmt.Println("Config Test failed: ", result)
		} else {
			fmt.Println("Config Test successful")
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
