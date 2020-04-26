package main

import (
	"fmt"

	"github.com/kabaluyot/gojson"
)

func main() {
	// Prints the JSON encoded string with keys preserved as specified
	om := gojson.OrderedMap{
		"name":  "Karl Anthony Baluyot",
		"age":   22,
		"likes": "smart girls",
	}

	fmt.Println("Sample 1", om.ToJSON("name", "age", "likes"))

	// Prints randomly all the keys
	fmt.Println("Sample 2", om.ToJSON())

	// Prints only the specified keys
	fmt.Println("Sample 2", om.ToJSON("likes"))
}
