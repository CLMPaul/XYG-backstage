package dberror

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

// @see: https://dev.mysql.com/doc/mysql-errors/5.7/en/server-error-reference.html
const (
	// Message: Duplicate entry '%s' for key %d
	ErrMySQLConstraintUnique uint16 = 1062
	// Message: Cannot add or update a child row: a foreign key constraint fails (%s)
	ErrMySQLConstraintForeignKey uint16 = 1452
)

func DowncastMySQLError(err error) (mysqlErr *mysql.MySQLError, ok bool) {
	ok = errors.As(err, &mysqlErr)
	return
}
