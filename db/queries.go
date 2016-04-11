package db

import (
	"database/sql"
	"fmt"
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

func BatchInsert(query BatchInsertQuery) (sql.Result, error) {
	columns, values := common.UnzipSlices(query.colvalues)
	insertColumns := strings.Join(columns, ", ")
	insertBindings := strings.Join(createBatchBindingsSlice(len(columns), len(values)), ", ")
	q := fmt.Sprintf("INSERT INTO %v (%v) VALUES %v", query.table, insertColumns, insertBindings)
	flatValues := flattenValues(values)
	return db.Exec(q, flatValues...)
}

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

func createBatchBindingsSlice(cols int, rows int) []string {
	batchBindings := make([]string, rows)
	for i := 0; i < rows; i++ {
		insertBindings := createBindingsSlice(cols*i+1, cols*(i+1))
		batchBindings[i] = fmt.Sprintf("(%v)", strings.Join(insertBindings, ", "))
	}
	return batchBindings
}

func flattenValues(values [][]interface{}) []interface{} {
	if len(values) == 0 {
		// TODO(nare469): find better way to handle
		return nil
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

func createBindingsSlice(start int, end int) []string {
	length := end - start + 1
	bindings := make([]string, length)
	for i := 0; i < length; i++ {
		bindings[i] = fmt.Sprintf("$%v", strconv.Itoa(start+i))
	}
	return bindings
}

func formatWheresAndOrderValues(wheres map[string]interface{}) (string, []interface{}) {
	columns, values := common.Unzip(wheres)
	whereClauses := make([]string, len(wheres))

	for i, column := range columns {
		whereClauses[i] = fmt.Sprintf("%v = $%v", column, i+1)
	}

	return strings.Join(whereClauses, " and "), values
}
