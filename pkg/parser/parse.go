package parser

import (
	"fmt"
	"reflect"

	"github.com/xqbumu/sqlparser"
)

func Parse(stmt sqlparser.Statement, level int) (any, error) {
	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		return parseSelect(stmt, level)

	default:
		return nil, fmt.Errorf("unsupported statement %+v of type %v", stmt, reflect.TypeOf(stmt))
	}
}

func parseSelect(stmt *sqlparser.Select, level int) (any, error) {
	if stmt.With != nil {
		parseWith(stmt.With, level)
	}

	return nil, nil
}

func parseWith(stmt *sqlparser.With, level int) (any, error) {
	sqlparser.VisitRefOfWith(stmt, func(node sqlparser.SQLNode) (kontinue bool, err error) {
		return false, nil
	})

	return nil, nil
}
