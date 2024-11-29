package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DBClient struct {
	db *sql.DB
}

// NewDBClient 构造函数，用于初始化 MySQL 数据库客户端
func NewDBClient(user, password, host string, port int, dbname string) (*DBClient, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	// 配置数据库连接池
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return &DBClient{db: db}, nil
}

// Query 示例方法：执行查询
func (c *DBClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return c.db.Query(query, args...)
}

// Close 关闭数据库连接
func (c *DBClient) Close() error {
	return c.db.Close()
}
