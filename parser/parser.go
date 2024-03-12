package parser

import (
	"github.com/tidwall/gjson"
)

// ParseJSONData takes a jsonData string and a path, queries the jsonData using the path,
// and returns the raw JSON string of the queried data.
func ParseJSONData(jsonData, path string) string {
	// Query the jsonData using the provided path
	result := gjson.Get(jsonData, path)

	// Return the raw JSON string of the result
	return result.Raw
}
