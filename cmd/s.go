/*
Copyright © 2023 Maneth PAK <manethpak00@gmail.com>
*/
package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

// sCmd represents the s command
var sCmd = &cobra.Command{
	Use:   "s",
	Short: "Search a query on google",
	Long:  "Search a query on google. For example: s how to google like a pro",
	Run: func(cmd *cobra.Command, args []string) {
		var search string

		if len(args) <= 0 && args[0] == "" {
			fmt.Println("Please provide a search term")
			return
		}
		search = fmt.Sprintf("%s site:(%s)", buildQuery(args), filterSite())
		search = url.QueryEscape(search)
		fmt.Println("https://google.com/search?q=" + search)
	},
}

func buildQuery(args []string) string {
	var query string

	if len(args) >= 1 && args[0] != "" {
		// search convert args to url query format
		query = strings.Join(args, " ")
	}

	return query
}

func filterSite() string {
	sites := []string{
		"stackoverflow.com",
		"github.com",
		"medium.com",
		"dev.to",
		"geeksforgeeks.org",
		"youtube.com",
	}

	return strings.Join(sites, " | ")
}

func init() {
	rootCmd.AddCommand(sCmd)
}
