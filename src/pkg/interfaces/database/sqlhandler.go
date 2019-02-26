package database

type SqlHandler interface {
	Exec(string, ...interface{}) error
	QueryRow(string, ...interface{}) (Row, error)
	Query(string, ...interface{}) (Rows, error)
}

type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

type Row interface {
	Scan(...interface{}) error
}
