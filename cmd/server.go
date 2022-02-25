/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"homework-1-ErdemOzgen/banner"
	"homework-1-ErdemOzgen/utils"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Booting up the restapi server and open browser to see the result",
	Long: `Function takes no parameters and boots up the restapi server also browser to see the result.
	you can use the following commands:
	
	=>          $ homework-1-ErdemOzgen server
	=>          $ homework-1-ErdemOzgen server -h
	=>          $ homework-1-ErdemOzgen server -help
	
	==> curl http://localhost:8000/books
	==> curl http://localhost:8000/books/author/Erdem
	==> curl http://localhost:8000/books/author/Erdem%20Ozgen
	==> curl http://localhost:8000/books/title/Golang%20Book
	`,
	Run: func(cmd *cobra.Command, args []string) {
		banner.PrintBanner()
		fmt.Println("server called")
		utils.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

}
