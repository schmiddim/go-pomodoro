/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

//go:embed .version
var version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "check version of command",
	Long:  `check version of command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
		fmt.Println("Go Version:", runtime.Version())
		fmt.Println("GOOS:", runtime.GOOS)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
