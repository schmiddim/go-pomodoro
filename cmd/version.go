/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var GitTag string

// GitCommit @see https://icinga.com/blog/2022/05/25/embedding-git-commit-information-in-go-binaries/
var GitCommit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}
	return ""
}()
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "check version of command",
	Long:  `check version of command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Code Version:", GitCommit)
		fmt.Println("Go Version:", runtime.Version())
		fmt.Println("Git Tag:", GitTag)
		fmt.Println("GOOS:", runtime.GOOS)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
