/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/nagaremono/tix/item"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todo items",
	Long:  `List todo items`,
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := item.ReadItems(viper.GetString("datafile"))
	if err != nil {
		fmt.Printf("%v", err)
	}

	sort.Sort(item.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 1, 1, ' ', 0)
	for _, i := range items {
		if allOpt || i.Done == doneOpt {
			fmt.Fprintln(w, i.PrettyDisplay())
		}
	}
	w.Flush()
}

var (
	doneOpt bool
	allOpt  bool
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' todo items")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all todo items")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
