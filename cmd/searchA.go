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
	Use:   "author",
	Short: "For searching by Authors' name",
	Long: `"func SearchByAuthor(author string, books []Book)" => function
	takes author name and books array as parameters Does not return anything.
	Just prints the result if exists in datasets.`,
	Run: func(cmd *cobra.Command, args []string) {
		banner.PrintBanner()
		//fmt.Println("searchA called")
		var s []string
		var books []search.Book

		s = append(s, args...)
		fmt.Println(strings.Join(s, " "))
		sd := strings.Join(args, " ")
		if sd == "" {
			fmt.Println("Please enter author's name")
			return
		}
		json.Unmarshal([]byte(jsonops.OpenJsonFile("test.json")), &books)

		search.SearchByAuthor(sd, books)
	},
}

func init() {
	rootCmd.AddCommand(searchACmd)

}
