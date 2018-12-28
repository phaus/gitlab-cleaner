package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	humanize "github.com/dustin/go-humanize"
	"github.com/phaus/gitlab-cleaner/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the content of a Registry.",
	Long:  `Lists the content of a Registry in sorted order.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := utils.GetClient()
		registries, err := utils.GetRegistry(client)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\nLISTING %s\n", viper.GetString("RegistryUrl"))
		for _, registry := range registries {
			registryTags, keys := utils.GetTagsAndSortedKeys(client, registry)
			fmt.Printf("| %4s | %20s | %20s | %11s | %10s | %s\n",
				"No",
				"Date",
				"Name",
				"Revision",
				"Size",
				"DELETE URL\n")
			for i, k := range keys {
				tagTime := utils.ParseTime(k)
				registryTag := registryTags[k]
				fmt.Printf("| %4d | %20s | %20s | %11s | %10s | DELETE %s%s\n",
					i+1,
					tagTime.Format("2006-01-02 15:04:05"),
					registryTag.Name,
					registryTag.ShortRevision,
					humanize.Bytes(registryTag.TotalSize),
					viper.GetString("BaseUrl"),
					registryTag.DestroyPath)
			}
		}
	},
}
