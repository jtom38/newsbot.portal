package api

import "strings"

// Converts the string of tags found in the database to a slice and cleaned up.
func splitTags(tags string) []string {
	var items []string
	
	temp := strings.Split(tags, ",")

	for _, i := range temp {
		i = strings.Trim(i, " ")
		items = append(items, i)
	}
	return items
}