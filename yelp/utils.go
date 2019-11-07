package yelp

import (
	"strings"
)

func CategoriesToString(cats []Category) string {

	result := []string{}

	for _, cat := range cats {
		if cat.Title != "" {
			result = append(result, cat.Title)
		}
	}

	joined := strings.Join(result, ",")

	return joined
}
