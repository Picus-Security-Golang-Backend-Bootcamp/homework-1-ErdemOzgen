/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "homework-1-ErdemOzgen",
	Short: "Homework for patika.dev Picus Bootcamp",
	Long: `This is project for patika.dev Picus Bootcamp homework 1.

#............^!JB@@#: .. J@@G5P55PB&@&PJ!^:....:~?5#@@@@@#^^^^~&@@@@@@@@@#~^^^~#@@@#5?~:.....:~7YB&@@
#               ~B@?    :&@#55PB&@G7:             .~5&@@B    .&@@@@@@@@@#.   .#@#!.              :G@
#    P#BBBBP!    .B&.    J@@B#@@#~    .!YGBBBG57:    :P@B    :&@@@@@@@@@#.   .#@~    !5GBBBG5?^ ^P@@
#    B@@@@@@@~    J@J    .#@@@@B.    ?#@@@@@@@@@&Y!?5G#@B    :&@@@@@@@@@#.   .#@:    J&@@@@@@@@#&@@@
#    B@@@@@&P.    P@&:    ?@@@@^    J@@@@@@@@@@@@@@@@@@@B    :&@@@@@@@@@#.   .#@5.     :^~7?Y5G#@@@@
#    ~!!!!~:     J@@@Y    .#@@#.    #@@@@@@@@@@@@@@@@@@@B    :&@@@@@@@@@#.   .#@@#Y!:.         .^?B@
#            .~JB@@@@&:    ?@@@^    Y@@@@@@@@@@@@@@@@@@@B    .&@@@@@@@@@#.   .#@@@@@&#GPYJ?!~.    .5
#    JP555PGB&@@@@@@@@Y    .#@@G.    ?&@@@@@@@@@&5!J5G#@&:    5@@@@@@@@@Y    ^@@@@P#@@@@@@@@@@J    :
#    B@@@@@@@@@@@@@@@@&:    7@@@B~    .!5GB##B57:    :P@@G.    !PB#&#B5!    .G@@G^ .~?5GB###B5^    !
#    B@@@@@@@@@@@@@@@@@5     B@@@@P!.      .      .^Y&@@@@#7.     ...     .7#@@&?:       ...     .?&
JJJJJ#@@@@@@@@@@@@@@@@@&J????G@@@@@@#P?~:.   .:^75B@@@@@@@@@#57~:.   .:~?5#@@@@@@&GY7~:.    .:^7Y#@@
			      PATIKA DEV BOOTCAMP ERDEM OZGEN
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.homework-1-ErdemOzgen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
