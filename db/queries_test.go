package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlattenValues(t *testing.T) {
	assert := assert.New(t)

	slices := [][]interface{}{{"a", "c"}, {"b", "d"}}
	expectedFlat := []interface{}{"a", "b", "c", "d"}

	flatValues := flattenValues(slices)
	assert.Equal(flatValues, expectedFlat)
}
