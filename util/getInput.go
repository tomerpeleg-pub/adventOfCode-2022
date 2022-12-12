package util

import (
	"log"
	"os"
)

func GetDayInput(day string) string {
	content, err := os.ReadFile("inputs/day" + day)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
