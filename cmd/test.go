package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/phaus/gitlab-cleaner/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dustin/go-humanize"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Tests the gitlab registry.",
	Long:  `This tests a gitlab Registry.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := utils.GetClient()
		registries, err := utils.GetRegistry(client)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\nTESTING %s\n", viper.GetString("RegistryUrl"))
		for _, registry := range registries {
			registryTags, err := utils.GetTags(client, registry)
			if err != nil {
				log.Fatal(err)
			}
			totalSize := countTotalSize(registryTags)
			fmt.Printf("%d %s, in total of %s created in %v.\n",
				len(registryTags),
				utils.ImageLabel(len(registryTags)),
				humanize.Bytes(totalSize),
				calculateDuration(registryTags))
			avgSize := totalSize / uint64(len(registryTags))
			fmt.Printf("The average image size is %s.\n", humanize.Bytes(avgSize))
		}
	},
}

func countTotalSize(registryTags map[string]utils.RegistryTag) uint64 {
	var count uint64
	for _, registryTag := range registryTags {
		count = count + registryTag.TotalSize
	}
	return count
}

func calculateDuration(registryTags map[string]utils.RegistryTag) time.Duration {
	keys := utils.SortedKeys(registryTags)
	start := utils.ParseTime(keys[0])
	end := utils.ParseTime(keys[len(keys)-1])
	return end.Sub(start)
}
