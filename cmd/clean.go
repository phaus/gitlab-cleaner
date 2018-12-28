package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/phaus/gitlab-cleaner/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var keepLast string
var keepDays string

func init() {
	rootCmd.AddCommand(cleanCmd)
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleans the content of a Registry.",
	Long:  `Cleans the content of a Registry based on parameters.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(keepDays) == 0 && len(keepLast) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		if len(keepDays) > 0 {
			days, err := strconv.Atoi(keepDays)
			if err != nil {
				log.Fatalf("Cannot parse %s to a number!", keepDays)
			}
			fmt.Printf("I keep all images of the last %d %s!\n", days, dayLabel(days))
			clean("days", days)
		}
		if len(keepLast) > 0 {
			last, err := strconv.Atoi(keepLast)
			if err != nil {
				log.Fatalf("Cannot parse %s to a number!", keepLast)
			}
			fmt.Printf("I will remove all images, but the last %d!\n", last)
			clean("last", last)
		}
	},
}

func init() {
	cleanCmd.Flags().StringVarP(&keepLast, "keep-last", "l", "", "Keep last X images.")
	cleanCmd.Flags().StringVarP(&keepDays, "keep-day", "k", "", "Keep images of the last X days.")
}

func dayLabel(count int) string {
	if count > 1 {
		return "days"
	}
	return "day"
}

func clean(filter string, value int) {
	client := utils.GetClient()
	registries, err := utils.GetRegistry(client)
	registryTags := make([]utils.RegistryTag, 0)

	if err != nil {
		log.Fatalf("Cannot load the registry")
	}
	if len(keepDays) > 0 {
		for _, registry := range registries {
			registryTags = collectTagsToDelete(client, registry, "days", value)
		}
	}
	if len(keepLast) > 0 {
		for _, registry := range registries {
			registryTags = collectTagsToDelete(client, registry, "last", value)
		}
	}
	deleteTags(client, registryTags)
}

func collectTagsToDelete(client *http.Client, registry utils.Registry, filter string, value int) []utils.RegistryTag {
	tags := make([]utils.RegistryTag, 0)
	registryTags, keys := utils.GetTagsAndSortedKeys(client, registry)

	if "days" == filter {
		now := time.Now()
		then := now.Add(time.Duration(-value) * time.Hour * 24)
		for _, k := range keys {
			tag := registryTags[k]
			tagTime := utils.ParseTime(k)
			if tagTime.Before(then) {
				tags = append(tags, tag)
			}
		}
	}

	if "last" == filter {
		for i, k := range keys {
			if i >= value {
				tags = append(tags, registryTags[k])
			}
		}
	}
	return tags
}

func deleteTags(client *http.Client, registryTags []utils.RegistryTag) {
	if DryRun {
		for _, tag := range registryTags {
			tagDestoryURL := fmt.Sprintf("%s%s", viper.GetString("BaseUrl"), tag.DestroyPath)
			log.Printf("[DRY] Deleting %s, %s", tag.Name, tagDestoryURL)
		}
	} else {
		for _, tag := range registryTags {
			tagDestoryURL := fmt.Sprintf("%s%s", viper.GetString("BaseUrl"), tag.DestroyPath)
			log.Printf("Deleting %s, %s", tag.Name, tagDestoryURL)
			err := utils.DeleteTag(client, tagDestoryURL)
			if err != nil {
				log.Printf("Error deleting %s: %s", tag.Name, err.Error())
			}
		}
	}
}

// ExtractDuration - extracts a duration from a parameter.
func ExtractDuration(parameter string) time.Duration {
	startingTime := time.Now().UTC()
	time.Sleep(10 * time.Millisecond)
	endingTime := time.Now().UTC()

	var duration = endingTime.Sub(startingTime)
	return duration
}
