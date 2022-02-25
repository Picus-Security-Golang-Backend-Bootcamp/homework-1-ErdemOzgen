/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"homework-1-ErdemOzgen/jsonops"
	"homework-1-ErdemOzgen/search"
	"strings"

	"github.com/spf13/cobra"
)

// searchTCmd represents the searchT command
var searchTCmd = &cobra.Command{
	Use:   "title",
	Short: "For searching by Authors' name",
	Long: `"func SearchByTitle(title string, books []Book)" => function
	takes books title and books array as parameters Does not return anything.
	Just prints the result if exists in datasets.
`,
	Run: func(cmd *cobra.Command, args []string) {

		//banner.PrintBanner()
		//fmt.Println("searchT called")
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

}
