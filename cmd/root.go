/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var countBytes bool
var countLines bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "runs a simplified version of wc",
	Long: `wc is a command line tool that parses text in files. 

This is made as a copy of said program for coding challenges.`,
	Run: func(cmd *cobra.Command, args []string) {
		if countBytes {
			file_name := args[0]

			file_data, err := os.ReadFile(file_name)

			if err != nil {
				panic(err)
			}

			fmt.Printf("%d %s\n", len(file_data), file_name)
		}

		if countLines {
			var line_count int
			file_name := args[0]

			file_data, err := os.ReadFile(file_name)

			if err != nil {
				panic(err)
			}

			scanner := bufio.NewScanner(bytes.NewReader(file_data))
			for scanner.Scan() {
				line_count++
			}

			fmt.Printf("%d %s\n", line_count, file_name)
		}
	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&countBytes, "bytes", "c", false, "Print the byte counts")
	rootCmd.Flags().BoolVarP(&countLines, "lines", "l", false, "Print the number of lines in a file.")
}
