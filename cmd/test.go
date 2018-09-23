package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/phaus/gitlab-cleaner/utils"
	"github.com/spf13/cobra"

	"github.com/dustin/go-humanize"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test the gitlab registry.",
	Long:  `This tests a gitlab Registry.`,
	Run: func(cmd *cobra.Command, args []string) {
		registries, err := utils.GetRegistry(GetClient())
		if err != nil {
			log.Fatal(err)
		}
		for _, registry := range registries {
			registryTags, err := utils.GetTags(GetClient(), registry)
			if err != nil {
				log.Fatal(err)
			}
			totalSize := countTotalSize(registryTags)
			fmt.Printf("%d Image have been created %s during %v.\n",
				len(registryTags),
				humanize.Bytes(totalSize),
				calculateDuration(registryTags))
			avgSize := totalSize / uint64(len(registryTags))
			fmt.Printf("Average image size is %s.\n", humanize.Bytes(avgSize))
		}
	},
}

func countTotalSize(registryTags map[time.Time]utils.RegistryTag) uint64 {
	var count uint64
	for _, registryTag := range registryTags {
		count = count + registryTag.TotalSize
	}
	return count
}

func calculateDuration(registryTags map[time.Time]utils.RegistryTag) time.Duration {
	keys := utils.SortedKeys(registryTags)
	start := utils.ParseTime(keys[0])
	end := utils.ParseTime(keys[len(keys)-1])
	return end.Sub(start)
}
