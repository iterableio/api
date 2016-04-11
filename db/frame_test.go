package db

import (
	"fmt"
	"testing"
)

func TestBatchBindingSlice(t *testing.T) {
	fmt.Println(createBatchBindingsSlice(2, 2))
}
