package actions

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

// construct find options based on sort and skip values
var createFindOneOptions = func(sort string, skip int64) *options.FindOneOptions {
	findOptions := options.FindOne()
	if len(sort) > 0 {
		sortMap := make(map[string]int)
		parts := strings.Split(sort, ",")
		for _, part := range parts {
			key := part
			direction := 1
			if part[0] == '+' || part[0] == ' ' {
				key = part[1:]
				direction = 1
			} else if part[0] == '-' {
				key = part[1:]
				direction = -1
			}
			if key == "id" {
				key = "_id"
			}
			sortMap[key] = direction
		}
		findOptions.SetSort(sort)
	} else {
		findOptions.SetSort(map[string]int{"_id": -1})
	}
	findOptions.SetSkip(skip)

	return findOptions
}