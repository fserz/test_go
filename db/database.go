package db

import "database/sql"

// Database 定义了数据库操作的接口
type Database interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Close() error
}
