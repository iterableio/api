package db

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/iterableio/api/common"
)

type BasicSelectQuery struct {
	table   string
	columns []string
	wheres  map[string]interface{}
}

type BasicInsertQuery struct {
	table     string
	colvalues map[string]interface{}
}

type BatchInsertQuery struct {
	table     string
	colvalues map[string][]interface{}
}

// Constructs colvalues using the `column` struct tag with only the given fields
func colvaluesFromModel(model interface{}, fields []string) (map[string]interface{}, error) {
	colvalues := make(map[string]interface{})
	structType := reflect.ValueOf(model)
	for _, field := range fields {
		typeField, found := structType.Type().FieldByName(field)
		if !found {
			return nil, errors.New(
				fmt.Sprintf("%v field does not exist", field),
			)
		}
		columnName := typeField.Tag.Get("column")
		if columnName == "" {
			return nil, errors.New(
				fmt.Sprintf("%v field does not have a column struct tag", typeField.Name),
			)
		}
		colvalues[columnName] = structType.FieldByName(field).Interface()
	}
	return colvalues, nil
}

// Inserts multiple rows into postgres
func BatchInsert(query BatchInsertQuery) (sql.Result, error) {
	columns, values := common.UnzipSlices(query.colvalues)
	insertColumns := strings.Join(columns, ", ")
	insertBindings := strings.Join(createBatchBindingsSlice(len(columns), len(values)), ", ")
	q := fmt.Sprintf("INSERT INTO %v (%v) VALUES %v", query.table, insertColumns, insertBindings)
	flatValues := flattenValues(values)
	return db.Exec(q, flatValues...)
}

// Inserts a single row into postgres
func Insert(query BasicInsertQuery) (sql.Result, error) {
	columns, values := common.Unzip(query.colvalues)
	insertColumns := strings.Join(columns, ", ")
	insertBindings := strings.Join(createBindingsSlice(1, len(columns)), ", ")
	q := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", query.table, insertColumns, insertBindings)
	return db.Exec(q, values...)
}

func Select(dest interface{}, query BasicSelectQuery) error {
	wheres, values := formatWheresAndOrderValues(query.wheres)
	selectColumns := strings.Join(query.columns, ", ")
	q := fmt.Sprintf("SELECT %v FROM %v WHERE %v", selectColumns, query.table, wheres)
	return db.Get(dest, q, values...)
}

// Returns a comma-separated list of bindings for a batch insert
func createBatchBindingsSlice(cols int, rows int) []string {
	batchBindings := make([]string, rows)
	for i := 0; i < rows; i++ {
		insertBindings := createBindingsSlice(cols*i+1, cols*(i+1))
		batchBindings[i] = fmt.Sprintf("(%v)", strings.Join(insertBindings, ", "))
	}
	return batchBindings
}

func createBindingsSlice(start int, end int) []string {
	length := end - start + 1
	bindings := make([]string, length)
	for i := 0; i < length; i++ {
		bindings[i] = fmt.Sprintf("$%v", strconv.Itoa(start+i))
	}
	return bindings
}

// Flattens a list of column values into an order that corresponds to the bindings
// slice returned from createBatchBindingsSlice
func flattenValues(values [][]interface{}) []interface{} {
	if len(values) == 0 {
		return []interface{}{}
	}
	numColumns := len(values)
	numRows := len(values[0])
	flatValues := make([]interface{}, numColumns*numRows)

	for col, _ := range values {
		for row, value := range values[col] {
			flatValues[row*numColumns+col] = value
		}
	}
	return flatValues
}

func formatWheresAndOrderValues(wheres map[string]interface{}) (string, []interface{}) {
	columns, values := common.Unzip(wheres)
	whereClauses := make([]string, len(wheres))

	for i, column := range columns {
		whereClauses[i] = fmt.Sprintf("%v = $%v", column, i+1)
	}

	return strings.Join(whereClauses, " and "), values
}
