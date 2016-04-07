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

func Insert(query BasicInsertQuery) (sql.Result, error) {
	columns, values := common.Unzip(query.colvalues)
	insertColumns := strings.Join(columns, ", ")

	bindings := make([]string, len(columns))
	for i := 0; i < len(bindings); i++ {
		bindings[i] = fmt.Sprintf("$%v", strconv.Itoa(i+1))
	}
	insertBindings := strings.Join(bindings, ", ")

	q := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", query.table, insertColumns, insertBindings)
	return db.Exec(q, values...)
}

func Select(dest interface{}, query BasicSelectQuery) error {
	wheres, values := formatWheresAndOrderValues(query.wheres)
	selectColumns := strings.Join(query.columns, ", ")
	q := fmt.Sprintf("SELECT %v FROM %v WHERE %v", selectColumns, query.table, wheres)
	return db.Get(dest, q, values...)
}

func formatWheresAndOrderValues(wheres map[string]interface{}) (string, []interface{}) {
	columns, values := common.Unzip(wheres)
	whereClauses := make([]string, len(wheres))

	for i, column := range columns {
		whereClauses[i] = fmt.Sprintf("%v = $%v", column, i+1)
	}

	return strings.Join(whereClauses, " and "), values
}
