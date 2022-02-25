/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"homework-1-ErdemOzgen/jsonops"
	"homework-1-ErdemOzgen/search"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all books in the dataset",
	Long: `"func ListBooks(books []Book)" => function takes book array 
	and prints out all books in the dataset.Does not return anything.`,
	Run: func(cmd *cobra.Command, args []string) {
		//banner.PrintBanner()
		fmt.Println("list called")
		var books []search.Book
		json.Unmarshal([]byte(jsonops.OpenJsonFile("test.json")), &books)
		search.ListBooks(books)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
