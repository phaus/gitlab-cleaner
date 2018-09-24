package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
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
		if len(keepDays) > 0 {
			days, err := strconv.Atoi(keepDays)
			if err != nil {
				log.Fatalf("Cannot parse %s to a number!", keepDays)
			}
			fmt.Printf("I keep all images of the last %s %s!\n", keepDays, dayLabel(days))
		}
		if len(keepLast) > 0 {
			fmt.Printf("I will remove all images, but the last %s!\n", keepLast)
		}
	},
}

func init() {
	cleanCmd.Flags().StringVarP(&keepLast, "keep-last", "l", "", "Keep last X images.")
	cleanCmd.Flags().StringVarP(&keepDays, "keep-day", "d", "", "Keep last images from X days.")
}

func dayLabel(count int) string {
	if count > 1 {
		return "days"
	}
	return "day"
}

// ExtractDuration - extracts a duration from a parameter.
func ExtractDuration(parameter string) time.Duration {
	startingTime := time.Now().UTC()
	time.Sleep(10 * time.Millisecond)
	endingTime := time.Now().UTC()

	var duration = endingTime.Sub(startingTime)
	return duration
}
