package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBClient struct {
	db *sql.DB
}

// 确保 DBClient 实现了 Database 接口
var _ Database = (*DBClient)(nil)

// NewDBClient 构造函数，用于初始化 MySQL 数据库客户端
func NewDBClient(user, password, host string, port int, dbname string) (*DBClient, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return &DBClient{db: db}, nil
}

// Query 方法，实现 Database 接口
func (c *DBClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return c.db.Query(query, args...)
}

// Close 方法，实现 Database 接口
func (c *DBClient) Close() error {
	return c.db.Close()
}
