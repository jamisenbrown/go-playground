/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/jamisenbrown/tri/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of the todo",
	Long:  `List will list all of the todos currently stored`,

	Run: func(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems(dataFile)

		if err != nil {
			log.Printf("%v", err)
		}

		fmt.Println(items)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
