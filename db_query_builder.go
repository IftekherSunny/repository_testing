package repository_testing

import "database/sql"

// Declaring variable
var dbQueryBuilder DbQueryBuilderContract

// Initiate dbQueryBuilder instance
func init() {
	dbQueryBuilder = new(DbQueryBuilder)
}

// DbQueryBuilder contract
type DbQueryBuilderContract interface {
	QueryRow(query string, args ...interface{})
	Scan(dest ...interface{}) error
}

// DbQueryBuilder struct
type DbQueryBuilder struct {
	row *sql.Row
}

// QueryRow executes a query that is expected to return at most one row
func (self *DbQueryBuilder) QueryRow(query string, args ...interface{}) {
	dbConnection := new(DbConnection).GetConnection()
	defer dbConnection.Close()

	self.row = dbConnection.QueryRow(query, args...)
}

// Scan copies the columns from the matched row into the values
func (self *DbQueryBuilder) Scan(dest ...interface{}) error {
	err := self.row.Scan(dest...)

	return err
}
