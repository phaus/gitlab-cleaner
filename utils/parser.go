package utils

import (
	"log"
	"time"
)

// ParseTime parses a string to a time object
func ParseTime(dateString string) time.Time {
	t, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
