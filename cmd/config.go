/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the future search filter for google operator",
	Long: `Configure the future search filter for google operator. use the --add-sites and --remove-sites flags to add or remove sites from the filter list.
For example: 'googlin config --add-sites stackoverflow.com,github.com' will add stackoverflow.com and github.com to the filter list.
Similarly, googlin config '--remove-sites stackoverflow.com,github.com' will remove stackoverflow.com and github.com from the filter list.`,
	Run: func(cmd *cobra.Command, args []string) {

		add, _ := cmd.Flags().GetString("add-sites")
		remove, _ := cmd.Flags().GetString("remove-sites")
		list, _ := cmd.Flags().GetBool("list")

		if add != "" {
			var newSites = strings.Split(add, ",")
			var sites = viper.GetStringSlice("sites")

			// check if the site is already present in the list
			for _, newSite := range newSites {
				var found = false
				for _, site := range sites {
					if site == newSite {
						found = true
						break
					}
				}
				if !found {
					sites = append(sites, newSite)
				}
			}

			viper.Set("sites", sites)
			viper.WriteConfig()

			log.Default().Println("Added sites:", strings.Join(newSites, ", "))
			return
		}

		if remove != "" {
			var removeSites = strings.Split(remove, ",")
			var sites = viper.GetStringSlice("sites")
			var newSites []string

			for _, site := range sites {
				var found = false
				for _, removeSite := range removeSites {
					if site == removeSite {
						found = true
						break
					}
				}
				if !found {
					newSites = append(newSites, site)
				}
			}
			viper.Set("sites", newSites)

			return
		}

		if list {
			var sites = viper.GetStringSlice("sites")
			log.Default().Println("All sites:")
			for _, site := range sites {
				log.Default().Println(" - " + site)
			}
			return
		}

		// check if any flag is present
		if add == "" && remove == "" && !list {
			log.Default().Println("Invalid command. Please use --help for more information.")
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	var addSites string
	var removeSites string

	configCmd.Flags().StringVarP(&addSites, "add-sites", "a", "", "Add filter the search results to the specified sites")
	configCmd.Flags().StringVarP(&removeSites, "remove-sites", "r", "", "Remove filter the search results to the specified sites")
	configCmd.Flags().BoolP("list", "l", false, "List all the sites in the filter list")
}
