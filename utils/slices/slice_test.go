package slices

import (
	"log"
	"testing"
)

func TestFilter(t *testing.T) {
	type obj struct {
		a int
		b int
	}

	input := []obj{
		{1, 2}, {3, 4}, {3, 5},
	}
	result := FirstObject(input, func(obj2 obj) bool { return obj2.a == 4 })
	log.Println(result)
}
