/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var flagDataSrc string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "f1-fantasy-teambuilder",
	Short: "f1-fantasy-teambuilder builds fantasy-teams for a given budget using simple statistical methods",
	Long: `A teambuilder for f1-fantasy teams based on simple statistical methods to find
optimal teams for any given budget. Built in GO.
Documentation available at https://github.com/asciiavenger/f1-fantasy-teambuilder`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.f1-fantasy-teambuilder.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.PersistentFlags().StringVarP(&flagDataSrc, "data", "d", "data.yaml", "YAML-File to read data from")
}
