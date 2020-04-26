package gojson

import (
	"fmt"
	"testing"
)

// Setting the om.ToJSON with the ordered keys will preserve its string encoded equivalent.
// Setting it to empty like om.toJSON() will randomly set its order just like the default map[string]interface to encoded equivalent.
func TestOrderedMap(t *testing.T) {
	om := OrderedMap{
		"name": "karl",
		"age":  21,
	}

	res := om.ToJSON("age")
	fmt.Println(res)
}

func BenchmarkOrderedMap(t *testing.B) {
	om := OrderedMap{
		"name": "karl",
		"age":  21,
	}

	res := om.ToJSON()
	fmt.Println(res)
}
