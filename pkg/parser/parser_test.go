package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xqbumu/sqlmatrix/pkg/parser"
	"github.com/xqbumu/sqlparser"
	. "github.com/xqbumu/sqlparser/examples"
)

func TestParseSimple(t *testing.T) {
	sql := "SELECT * FROM `table` WHERE a = 'abc'"
	stmt, err := sqlparser.Parse(sql)
	assert.Nil(t, err)
	assert.NotNil(t, stmt)

	result, err := parser.Parse(stmt, 0)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestParseWith(t *testing.T) {
	sql, err := FS.ReadFile("with_001.sql")
	assert.Nil(t, err)

	stmt, err := sqlparser.Parse(string(sql))
	assert.Nil(t, err)
	assert.NotNil(t, stmt)

	result, err := parser.Parse(stmt, 0)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
