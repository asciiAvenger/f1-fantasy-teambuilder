/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"f1-fantasy-teambuilder/printer"
	"f1-fantasy-teambuilder/ranker"
	"f1-fantasy-teambuilder/reader"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var driversCmd = &cobra.Command{
	Use:   "drivers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		data := reader.CollectData()
		drivers := ranker.RankDrivers(data, flagReverseOutput)
		printer.PrintRankResults(drivers)
	},
}

func init() {
	rankCmd.AddCommand(driversCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
