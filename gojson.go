/*
	Custom JSON parser using orderedmap

	@author: Karl Anthony Baluyot
	License: MIT

	See gojson_test.go for see testing.
*/

package gojson

import (
	"bytes"
	"fmt"
)

//OrderedMap type for ordered Json conversion
type OrderedMap map[string]interface{}

//ToJSON converts an OrderedMap type to Json string
func (om OrderedMap) ToJSON(order ...string) string {
	//check if no order param given
	if order == nil {
		for key := range om {
			order = append(order, key)
		}
	}

	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	l := len(order)

	for i, k := range order {
		fmt.Fprintf(buf, "\"%s\":\"%v\"", k, om[k])
		if i != (l - 1) {
			buf.WriteByte(',')
		}
	}
	buf.Write([]byte{'}'})

	return buf.String()
}
