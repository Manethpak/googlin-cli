/*
Copyright © 2023 Maneth PAK <manethpak00@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sCmd represents the s command
var sCmd = &cobra.Command{
	Use:   "s",
	Short: "Search a query on google",
	Long:  "Search a query on google. For example: s how to google like a pro",
	Run: func(cmd *cobra.Command, args []string) {
		var search string

		if len(args) <= 0 {
			log.Default().Println("Please provide a search term")
			return
		}
		search = fmt.Sprintf("%s site:(%s)", buildQuery(args), filterSite())
		search = url.QueryEscape(search)
		openbrowser(fmt.Sprintf("https://www.google.com/search?q=%s", search))
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
	var sites = viper.GetStringSlice("sites")
	return strings.Join(sites, " | ")
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	rootCmd.AddCommand(sCmd)
}
