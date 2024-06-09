package cmd

import (
	"fmt"
	"github.com/schmiddim/go-pomodoro/libs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

func countdown(minutes int, done chan bool) {
	for i := minutes * 60; i >= 0; i-- {
		fmt.Printf("\rTime remaining: %02d:%02d", i/60, i%60)
		time.Sleep(1 * time.Second)
	}
	done <- true
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-pomodoro",
	Short: "A Pomodoro timer CLI application",
	Long: `go-pomodoro is a command-line application that helps you manage your time using the Pomodoro Technique. 
You can set the duration of your work intervals and the application will notify you when it's time to take a break. 
It also integrates with an API to track your accomplishments.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		minutes, _ := cmd.Flags().GetInt("time")
		fmt.Printf("Pomodoro started. It will end in %d minutes.\n", minutes)
		timer := time.NewTimer(time.Duration(minutes) * time.Minute)
		done := make(chan bool)
		go countdown(minutes, done)

		<-timer.C
		<-done
		fmt.Println("Pomodoro ended.")

		r := libs.NewRest()
		err := r.SendAccomplishmentRequest(13)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	viper.SetDefault("POMODORO_API_KEY", "MeinSuperDuperApiKey")
	viper.SetDefault("POMODORO_ENDPOINT", "http://localhost:8080/api/accomplishments")
	viper.AutomaticEnv()
	rootCmd.PersistentFlags().IntP("time", "t", 30, "Set the timer duration in minutes")
}
