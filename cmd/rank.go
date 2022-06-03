/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

var flagReverseOutput bool

// buildCmd represents the build command
var rankCmd = &cobra.Command{
	Use:   "rank",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		// data := reader.CollectData()
	// 		// results := ranker.RankDrivers
	// 		// printer.PrintRankResults(results, flagTopTeams, flagBudget)
	// 	},
}

func init() {
	rootCmd.AddCommand(rankCmd)
	rankCmd.PersistentFlags().BoolVarP(&flagReverseOutput, "reverse", "r", false, "Reverses the output order")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
