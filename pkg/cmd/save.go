/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	internal "github.com/jordabrell/swap/pkg"
	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save the actual configuration of ~/.aws/credentials",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		internal.SaveConfiguration()
		fmt.Println("swap: saved well!")
		fmt.Println("swap: your configuration is saved in ~/.swap/saved-configuration")
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
