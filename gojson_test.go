package gojson

import (
	"fmt"
	"testing"
)

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

	res := om.ToJSON("age")
	fmt.Println(res)
}
