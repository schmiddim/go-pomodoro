package cmd

import (
	"fmt"
	"github.com/schmiddim/go-pomodoro/libs"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print Config",
	Long:  `Set config with Environment Variables POMODORO_API_KEY and POMODORO_ENDPOINT`,
	Run: func(cmd *cobra.Command, args []string) {

		r := libs.NewRest()

		fmt.Println(r)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
