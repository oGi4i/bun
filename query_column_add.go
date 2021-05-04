package bun

import (
	"fmt"

	"github.com/uptrace/bun/sqlfmt"
)

type AddColumnQuery struct {
	baseQuery
}

func NewAddColumnQuery(db *DB) *AddColumnQuery {
	q := &AddColumnQuery{
		baseQuery: baseQuery{
			db:  db,
			dbi: db.DB,
		},
	}
	return q
}

func (q *AddColumnQuery) DB(db DBI) *AddColumnQuery {
	q.dbi = db
	return q
}

func (q *AddColumnQuery) Model(model interface{}) *AddColumnQuery {
	q.setTableModel(model)
	return q
}

//------------------------------------------------------------------------------

func (q *AddColumnQuery) Table(tables ...string) *AddColumnQuery {
	for _, table := range tables {
		q.addTable(sqlfmt.UnsafeIdent(table))
	}
	return q
}

func (q *AddColumnQuery) TableExpr(query string, args ...interface{}) *AddColumnQuery {
	q.addTable(sqlfmt.SafeQuery(query, args))
	return q
}

func (q *AddColumnQuery) ModelTableExpr(query string, args ...interface{}) *AddColumnQuery {
	q.modelTable = sqlfmt.SafeQuery(query, args)
	return q
}

func (q *AddColumnQuery) ColumnExpr(query string, args ...interface{}) *AddColumnQuery {
	q.addColumn(sqlfmt.SafeQuery(query, args))
	return q
}

//------------------------------------------------------------------------------

func (q *AddColumnQuery) AppendQuery(fmter sqlfmt.QueryFormatter, b []byte) (_ []byte, err error) {
	if q.err != nil {
		return nil, q.err
	}
	if len(q.columns) != 1 {
		return nil, fmt.Errorf("bun: AddColumnQuery requires exactly one column")
	}

	b = append(b, "ALTER TABLE "...)

	b, err = q.appendFirstTable(fmter, b)
	if err != nil {
		return nil, err
	}

	b = append(b, " ADD "...)

	b, err = q.columns[0].AppendQuery(fmter, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}