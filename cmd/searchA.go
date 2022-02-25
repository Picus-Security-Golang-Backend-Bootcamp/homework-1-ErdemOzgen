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

// searchACmd represents the searchA command
var searchACmd = &cobra.Command{
	Use:   "searchA",
	Short: "Search by Author",
	Long: `func SearchByAuthor(author string, books []Book) function
	takes author name and books array as parameters Does not return anything.
	Just prints the result.`,
	Run: func(cmd *cobra.Command, args []string) {
		banner.PrintBanner()
		fmt.Println("searchA called")
		var s []string
		var books []search.Book

		s = append(s, args...)
		fmt.Println(strings.Join(s, " "))
		sd := strings.Join(args, " ")

		json.Unmarshal([]byte(jsonops.OpenJsonFile("test.json")), &books)

		search.SearchByAuthor(sd, books)
	},
}

func init() {
	rootCmd.AddCommand(searchACmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchACmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchACmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
