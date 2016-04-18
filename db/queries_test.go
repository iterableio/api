package db

import (
	"errors"
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

func TestFlattenValuesEmpty(t *testing.T) {
	assert := assert.New(t)

	emptyList := [][]interface{}{}
	expectedFlat := []interface{}{}

	flatValues := flattenValues(emptyList)
	assert.Equal(flatValues, expectedFlat)
}

func TestCreateBindingsSlice(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(createBindingsSlice(1, 3), []string{"$1", "$2", "$3"})
}

func TestCreateBatchBindingsSlice(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(createBatchBindingsSlice(2, 3), []string{"($1, $2)", "($3, $4)", "($5, $6)"})
}

func TestColvaluesFromModel(t *testing.T) {
	assert := assert.New(t)

	type testStruct struct {
		Name   string `column:"test_name"`
		Number int    `column:"test_number"`
	}

	type testWrongStruct struct {
		NoStructTag string
	}

	cases := []struct {
		model             interface{}
		fields            []string
		expectedColvalues map[string]interface{}
		expectedErr       error
	}{
		{
			model: testStruct{
				Name:   "Ayy lmao",
				Number: 4,
			},
			fields: []string{"Name", "Number"},
			expectedColvalues: map[string]interface{}{
				"test_name":   "Ayy lmao",
				"test_number": 4,
			},
			expectedErr: nil,
		},
		{
			model: testStruct{
				Name:   "Ayy lmao",
				Number: 4,
			},
			fields: []string{"Name"},
			expectedColvalues: map[string]interface{}{
				"test_name": "Ayy lmao",
			},
			expectedErr: nil,
		},
		{
			model: testStruct{
				Name:   "Ayy lmao",
				Number: 4,
			},
			fields:            []string{"IncorrectField"},
			expectedColvalues: nil,
			expectedErr:       errors.New("IncorrectField field does not exist"),
		},
		{
			model: testWrongStruct{
				NoStructTag: "Ayy lmao",
			},
			fields:            []string{"NoStructTag"},
			expectedColvalues: nil,
			expectedErr:       errors.New("NoStructTag field does not have a column struct tag"),
		},
	}
	for _, c := range cases {
		colvalues, err := colvaluesFromModel(c.model, c.fields)
		assert.Equal(colvalues, c.expectedColvalues)
		assert.Equal(err, c.expectedErr)
	}
}
