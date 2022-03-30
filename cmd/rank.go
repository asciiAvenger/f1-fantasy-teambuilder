/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"f1-fantasy-teambuilder/builder"
	"f1-fantasy-teambuilder/reader"

	"github.com/spf13/cobra"
)

// rankCmd represents the rank command
var rankCmd = &cobra.Command{
	Use:   "rank",
	Short: "Calculates scores and ranks drivers",
	Long: `Reads drivers stats from YAML-File and ranks them based
on a score that is calculated by averaging their performance
in fantasy-points over their cost.`,
	Run: func(cmd *cobra.Command, args []string) {
		data := reader.ReadData(flagDataSrc)
		builder.CreateRanking(data)
	},
}

func init() {
	rootCmd.AddCommand(rankCmd)

	// rankCmd.Flags().StringVarP(&dataSrc, "data", "d", "data.yaml", "YAML-File to read data from")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rankCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rankCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
