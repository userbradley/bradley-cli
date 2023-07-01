/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package gcloud

import (
	"fmt"
	"github.com/spf13/cobra"
)

// gcloudCmd represents the gcloud command
var Gcloud = &cobra.Command{
	Use:   "gcloud",
	Short: "Commands to help you interact with Google CLoud",
	Long:  `A selection of carefully curated commands that help you with your day to day interactions with Google cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello!")
	},
}

func init() {

	Gcloud.AddCommand(numberCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gcloudCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gcloudCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
