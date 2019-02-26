package infra

import (
	"database/sql"

	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/interfaces/database"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() database.SqlHandler {
	conn, err := sql.Open("mysql", "root:@tcp(db:3306)/sample")
	if err != nil {
		panic(err.Error)
	}
	return &SqlHandler{Conn: conn}
}

func (handler *SqlHandler) Exec(statement string, args ...interface{}) error {
	_, err := handler.Conn.Exec(statement, args...)
	return err
}

func (handler *SqlHandler) QueryRow(statement string, args ...interface{}) database.Row {
	row := handler.Conn.QueryRow(statement, args...)

	sRow := SqlRow{}
	sRow.Row = row
	return sRow
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Rows, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return SqlRows{}, err
	}
	sRows := SqlRows{}
	sRows.Rows = rows
	return sRows, nil
}

type SqlRow struct {
	Row *sql.Row
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}

type SqlRows struct {
	Rows *sql.Rows
}

func (r SqlRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRows) Next() bool {
	return r.Rows.Next()
}

func (r SqlRows) Close() error {
	return r.Rows.Close()
}
