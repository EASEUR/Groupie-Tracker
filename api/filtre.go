package api

import (
	"strconv"
	"strings"
)

func GetYear(date string) int {
	if len(date) < 4 {
		return 0
	}
	year, _ := strconv.Atoi(date[:4])
	return year
}

func ContainsLocation(locations []string, selected []string) bool {
	for _, l := range locations {
		for _, s := range selected {
			if strings.Contains(strings.ToLower(l), strings.ToLower(s)) {
				return true
			}
		}
	}
	return false
}
