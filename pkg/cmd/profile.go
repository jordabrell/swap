/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	internal "github.com/jordabrell/swap/pkg"

	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Swap the Default profile that you want.",
	Long: `Swap the Default profile that you want.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		
		profileName := args[0]
		internal.CheckArray(profileName)
		fmt.Println("Your default profile is: ",profileName)
		internal.ChangeProfile(profileName)
		internal.DeleteBridge()
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// profileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
