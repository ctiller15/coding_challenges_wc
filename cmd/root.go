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
var countWords bool
var countChars bool

func displayBytesCount(file_name string) int {
	file_data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}

	return len(file_data)
}

func displayLinesCount(file_name string) int {
	var line_count int

	file_data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(file_data))
	for scanner.Scan() {
		line_count++
	}

	return line_count
}

func displayWordsCount(file_name string) int {
	var word_count int

	file_data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(file_data))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word_count++
	}

	return word_count
}

func displayCharsCount(file_name string) int {
	var char_count int

	file_data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(file_data))
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		char_count++
	}

	return char_count
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "runs a simplified version of wc",
	Long: `wc is a command line tool that parses text in files. 

This is made as a copy of said program for coding challenges.`,
	Run: func(cmd *cobra.Command, args []string) {
		file_name := args[0]
		if countBytes {
			count := displayBytesCount(file_name)

			fmt.Printf("%d %s\n", count, file_name)
		}

		if countLines {
			count := displayLinesCount(file_name)

			fmt.Printf("%d %s\n", count, file_name)
		}

		if countWords {
			count := displayWordsCount(file_name)

			fmt.Printf("%d %s\n", count, file_name)
		}

		if countChars {
			count := displayCharsCount(file_name)

			fmt.Printf("%d %s\n", count, file_name)
		}

		if !countBytes && !countWords && !countLines && !countChars {
			word_count := displayWordsCount(file_name)
			line_count := displayLinesCount(file_name)
			byte_count := displayBytesCount(file_name)

			fmt.Printf("%d %d %d %s\n", line_count, word_count, byte_count, file_name)
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
	rootCmd.Flags().BoolVarP(&countWords, "words", "w", false, "Print the number of words in a file.")
	rootCmd.Flags().BoolVarP(&countChars, "chars", "m", false, "Print the number of chars in a file.")
}
