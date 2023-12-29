package utils

import (
	"strconv"
)

const ISO8601 string = `2006-01-02T15:04:05.999Z`

func IntFromQuery(query map[string]string, queryString string, defaultInt int) int {
	var integer int
	if query[queryString] != "" {
		i, err := strconv.ParseInt(query[queryString], 10, 64)
		integer = int(i)
		if err != nil {
			return 20 // TODO: Limit had a string inside
		}
	} else {
		integer = defaultInt
	}
	return integer
}
