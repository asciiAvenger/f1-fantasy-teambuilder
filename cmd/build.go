/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"f1-fantasy-teambuilder/builder"
	"f1-fantasy-teambuilder/printer"
	"f1-fantasy-teambuilder/reader"

	"github.com/spf13/cobra"
)

var flagBudget float64
var flagTopTeams int

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		data := reader.CollectData()
		results := builder.BuildTeam(data, flagBudget, flagTopTeams)
		printer.PrintBuildResults(results, flagTopTeams, flagBudget)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().Float64VarP(&flagBudget, "budget", "b", 100.0, "Maximum budget for the team im M$")
	buildCmd.Flags().IntVarP(&flagTopTeams, "top", "t", 5, "Top <x> teams to print out")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
