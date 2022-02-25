/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"homework-1-ErdemOzgen/banner"
	"homework-1-ErdemOzgen/jsonops"
	"homework-1-ErdemOzgen/search"
	"strings"

	"github.com/spf13/cobra"
)

// searchTCmd represents the searchT command
var searchTCmd = &cobra.Command{
	Use:   "searchT",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		banner.PrintBanner()
		fmt.Println("searchT called")
		var s []string
		var books []search.Book

		s = append(s, args...)
		fmt.Println(strings.Join(s, " "))
		sd := strings.Join(args, " ")
		if sd == "" {
			fmt.Println("Please enter book's title")
			return
		}
		json.Unmarshal([]byte(jsonops.OpenJsonFile("test.json")), &books)
		search.SearchByTitle(sd, books)
	},
}

func init() {
	rootCmd.AddCommand(searchTCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchTCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchTCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
