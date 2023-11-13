/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/jamisenbrown/tri/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark Item as done",
	Run:     doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if i > 0 && i < len(items) {
		items[i-1].Done = true
		fmt.Printf("%w %v\n", items[i-1].Text, "marked done")

		sort.Sort(todo.ByPri(items))
		todo.SaveItems(dataFile, items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}
