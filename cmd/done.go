/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/nagaremono/tix/item"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "marks an item as done",
	Long:  `marks an item as done`,
	Run:   doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := item.ReadItems(viper.GetString("datafile"))

	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if i > 0 && i < len(items) {
		items[i-1].Done = true

		fmt.Printf("%q %v\n", items[i-1].Content, "ticked done")

		sort.Sort(item.ByPri(items))
		item.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
