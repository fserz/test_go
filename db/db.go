package db

import (
	"fmt"
	"log"
)

func main() {
	// 使用 Database 接口来接收 DBClient 实例
	var db Database
	db, err := NewDBClient("root", "password", "localhost", 3306, "mydatabase")
	if err != nil {
		log.Fatalf("Failed to create DB client: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM users WHERE active = ?", true)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		fmt.Printf("User: %d, %s\n", id, name)
	}
}
